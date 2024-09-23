// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"context"
	"database/sql"
	"io"

	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/cloud"
	"github.com/juju/juju/controller"
	"github.com/juju/juju/core/credential"
	coremodel "github.com/juju/juju/core/model"
	modeltesting "github.com/juju/juju/core/model/testing"
	coreobjectstore "github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/core/permission"
	"github.com/juju/juju/core/providertracker"
	coreuser "github.com/juju/juju/core/user"
	jujuversion "github.com/juju/juju/core/version"
	userbootstrap "github.com/juju/juju/domain/access/bootstrap"
	cloudbootstrap "github.com/juju/juju/domain/cloud/bootstrap"
	cloudstate "github.com/juju/juju/domain/cloud/state"
	controllerconfigbootstrap "github.com/juju/juju/domain/controllerconfig/bootstrap"
	credentialbootstrap "github.com/juju/juju/domain/credential/bootstrap"
	modeldomain "github.com/juju/juju/domain/model"
	modelbootstrap "github.com/juju/juju/domain/model/bootstrap"
	modelconfigbootstrap "github.com/juju/juju/domain/modelconfig/bootstrap"
	modeldefaultsbootstrap "github.com/juju/juju/domain/modeldefaults/bootstrap"
	schematesting "github.com/juju/juju/domain/schema/testing"
	backendbootstrap "github.com/juju/juju/domain/secretbackend/bootstrap"
	domainservicefactory "github.com/juju/juju/domain/servicefactory"
	"github.com/juju/juju/internal/auth"
	databasetesting "github.com/juju/juju/internal/database/testing"
	loggertesting "github.com/juju/juju/internal/logger/testing"
	"github.com/juju/juju/internal/servicefactory"
	jujutesting "github.com/juju/juju/internal/testing"
	"github.com/juju/juju/internal/uuid"
)

// ServiceFactorySuite is a test suite that can be composed into tests that
// require a Juju ServiceFactory and database access. It holds the notion of a
// controller model uuid and that of a default model uuid. Both of these models
// will be instantiated into the database upon test setup.
type ServiceFactorySuite struct {
	schematesting.ControllerModelSuite

	// AdminUserUUID is the uuid of the admin user made during the setup of this
	// test suite.
	AdminUserUUID coreuser.UUID

	// CloudName is the name of the cloud made during the setup of this suite.
	CloudName string

	CredentialKey credential.Key

	// ControllerModelUUID is the unique id for the controller model. If not set
	// will be set during test set up.
	ControllerModelUUID coremodel.UUID

	// ControllerConfig is the controller configuration, including its UUID. If
	// not set will be set to the default testing value during test set up.
	ControllerConfig controller.Config

	// DefaultModelUUID is the unique id for the default model. If not set
	// will be set during test set up.
	DefaultModelUUID coremodel.UUID

	// ProviderTracker is the provider tracker to use in the service factory.
	ProviderTracker providertracker.ProviderFactory
}

type stubDBDeleter struct {
	DB *sql.DB
}

func (s stubDBDeleter) DeleteDB(namespace string) error {
	return nil
}

// ControllerServiceFactory conveniently constructs a service factory for the
// controller model.
func (s *ServiceFactorySuite) ControllerServiceFactory(c *gc.C) servicefactory.ServiceFactory {
	return s.ServiceFactoryGetter(c, TestingObjectStore{})(s.ControllerModelUUID)
}

// DefaultModelServiceFactory conveniently constructs a service factory for the
// default model.
func (s *ServiceFactorySuite) DefaultModelServiceFactory(c *gc.C) servicefactory.ServiceFactory {
	return s.ServiceFactoryGetter(c, TestingObjectStore{})(s.ControllerModelUUID)
}

func (s *ServiceFactorySuite) SeedControllerConfig(c *gc.C) {
	fn := controllerconfigbootstrap.InsertInitialControllerConfig(
		s.ControllerConfig,
		s.ControllerModelUUID,
	)
	err := fn(context.Background(), s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)
}

func (s *ServiceFactorySuite) SeedAdminUser(c *gc.C) {
	password := auth.NewPassword("dummy-secret")
	uuid, fn := userbootstrap.AddUserWithPassword(
		coreuser.AdminUserName,
		password,
		permission.AccessSpec{
			Access: permission.SuperuserAccess,
			Target: permission.ID{
				ObjectType: permission.Controller,
				Key:        jujutesting.ControllerTag.Id(),
			},
		},
	)
	s.AdminUserUUID = uuid
	err := fn(context.Background(), s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)
}

func (s *ServiceFactorySuite) SeedCloudAndCredential(c *gc.C) {
	ctx := context.Background()

	err := cloudstate.AllowCloudType(ctx, s.ControllerTxnRunner(), 99, "dummy")
	c.Assert(err, jc.ErrorIsNil)

	s.CloudName = "dummy"
	err = cloudbootstrap.InsertCloud(coreuser.AdminUserName, cloud.Cloud{
		Name:      s.CloudName,
		Type:      "dummy",
		AuthTypes: []cloud.AuthType{cloud.EmptyAuthType, cloud.AccessKeyAuthType, cloud.UserPassAuthType},
		Regions: []cloud.Region{
			{
				Name: "dummy-region",
			},
		},
	})(ctx, s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)

	s.CredentialKey = credential.Key{
		Cloud: s.CloudName,
		Name:  "default",
		Owner: coreuser.AdminUserName,
	}
	err = credentialbootstrap.InsertCredential(
		s.CredentialKey,
		cloud.NewCredential(cloud.UserPassAuthType, map[string]string{
			"username": "dummy",
			"password": "secret",
		}),
	)(ctx, s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)
}

// SeedModelDatabases makes sure that model's for both the controller and default
// model have been created in the database.
func (s *ServiceFactorySuite) SeedModelDatabases(c *gc.C) {
	ctx := context.Background()

	controllerUUID, err := uuid.UUIDFromString(jujutesting.ControllerTag.Id())
	c.Assert(err, jc.ErrorIsNil)

	controllerArgs := modeldomain.ModelCreationArgs{
		AgentVersion: jujuversion.Current,
		Cloud:        s.CloudName,
		CloudRegion:  "dummy-region",
		Credential:   s.CredentialKey,
		Name:         coremodel.ControllerModelName,
		Owner:        s.AdminUserUUID,
	}

	fn := modelbootstrap.CreateModel(s.ControllerModelUUID, controllerArgs)
	c.Assert(backendbootstrap.CreateDefaultBackends(coremodel.IAAS)(
		ctx, s.ControllerTxnRunner(), s.ModelTxnRunner(c, s.ControllerModelUUID.String())), jc.ErrorIsNil)
	err = fn(ctx, s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)

	err = modelbootstrap.CreateReadOnlyModel(s.ControllerModelUUID, controllerUUID)(ctx, s.ControllerTxnRunner(), s.ModelTxnRunner(c, s.ControllerModelUUID.String()))
	c.Assert(err, jc.ErrorIsNil)

	fn = modelconfigbootstrap.SetModelConfig(
		s.ControllerModelUUID,
		nil,
		modeldefaultsbootstrap.ModelDefaultsProvider(nil, nil),
	)
	err = fn(ctx, s.ControllerTxnRunner(), s.ModelTxnRunner(c, s.ControllerModelUUID.String()))
	c.Assert(err, jc.ErrorIsNil)

	modelArgs := modeldomain.ModelCreationArgs{
		AgentVersion: jujuversion.Current,
		Cloud:        s.CloudName,
		Credential:   s.CredentialKey,
		Name:         "test",
		Owner:        s.AdminUserUUID,
	}

	fn = modelbootstrap.CreateModel(s.DefaultModelUUID, modelArgs)
	err = fn(ctx, s.ControllerTxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)

	err = modelbootstrap.CreateReadOnlyModel(s.DefaultModelUUID, controllerUUID)(ctx, s.ControllerTxnRunner(), s.ModelTxnRunner(c, s.DefaultModelUUID.String()))
	c.Assert(err, jc.ErrorIsNil)

	fn = modelconfigbootstrap.SetModelConfig(
		s.DefaultModelUUID,
		nil,
		modeldefaultsbootstrap.ModelDefaultsProvider(nil, nil),
	)
	err = fn(ctx, s.ControllerTxnRunner(), s.ModelTxnRunner(c, s.DefaultModelUUID.String()))
	c.Assert(err, jc.ErrorIsNil)
}

// ServiceFactoryGetter provides an implementation of the ServiceFactoryGetter
// interface to use in tests.
func (s *ServiceFactorySuite) ServiceFactoryGetter(c *gc.C, objectStore coreobjectstore.ObjectStore) ServiceFactoryGetterFunc {
	return func(modelUUID coremodel.UUID) servicefactory.ServiceFactory {
		return domainservicefactory.NewServiceFactory(
			databasetesting.ConstFactory(s.TxnRunner()),
			modelUUID,
			databasetesting.ConstFactory(s.ModelTxnRunner(c, modelUUID.String())),
			stubDBDeleter{DB: s.DB()},
			s.ProviderTracker,
			singularObjectStoreGetter(func(ctx context.Context) (coreobjectstore.ObjectStore, error) {
				return objectStore, nil
			}),
			loggertesting.WrapCheckLog(c),
		)
	}
}

// ObjectStoreServicesGetter provides an implementation of the
// ObjectStoreServicesGetter interface to use in tests.
func (s *ServiceFactorySuite) ObjectStoreServicesGetter(c *gc.C) ObjectStoreServicesGetterFunc {
	return func(modelUUID coremodel.UUID) servicefactory.ObjectStoreServices {
		return domainservicefactory.NewObjectStoreServices(
			databasetesting.ConstFactory(s.TxnRunner()),
			databasetesting.ConstFactory(s.ModelTxnRunner(c, modelUUID.String())),
			loggertesting.WrapCheckLog(c),
		)
	}
}

// NoopObjectStore returns a no-op implementation of the ObjectStore interface.
// This is useful when the test does not require any object store functionality.
func (s *ServiceFactorySuite) NoopObjectStore(c *gc.C) coreobjectstore.ObjectStore {
	return TestingObjectStore{}
}

// SetUpTest creates the controller and default model unique identifiers if they
// have not already been set. Also seeds the initial database with the models.
func (s *ServiceFactorySuite) SetUpTest(c *gc.C) {
	s.ControllerModelSuite.SetUpTest(c)
	if s.ControllerModelUUID == "" {
		s.ControllerModelUUID = modeltesting.GenModelUUID(c)
	}
	if s.ControllerConfig == nil {
		s.ControllerConfig = jujutesting.FakeControllerConfig()
	}
	if s.DefaultModelUUID == "" {
		s.DefaultModelUUID = modeltesting.GenModelUUID(c)
	}
	s.SeedControllerConfig(c)
	s.SeedAdminUser(c)
	s.SeedCloudAndCredential(c)
	s.SeedModelDatabases(c)
}

// ServiceFactoryGetterFunc is a convenience type for translating a getter
// function into the ServiceFactoryGetter interface.
type ServiceFactoryGetterFunc func(coremodel.UUID) servicefactory.ServiceFactory

// FactoryForModel implements the ServiceFactoryGetter interface.
func (s ServiceFactoryGetterFunc) FactoryForModel(modelUUID coremodel.UUID) servicefactory.ServiceFactory {
	return s(modelUUID)
}

// ObjectStoreServicesGetterFunc is a convenience type for translating a getter
// function into the ObjectStoreServicesGetter interface.
type ObjectStoreServicesGetterFunc func(coremodel.UUID) servicefactory.ObjectStoreServices

// FactoryForModel implements the ObjectStoreServicesGetter interface.
func (s ObjectStoreServicesGetterFunc) FactoryForModel(modelUUID coremodel.UUID) servicefactory.ObjectStoreServices {
	return s(modelUUID)
}

type singularObjectStoreGetter func(context.Context) (coreobjectstore.ObjectStore, error)

func (s singularObjectStoreGetter) GetObjectStore(ctx context.Context) (coreobjectstore.ObjectStore, error) {
	return s(ctx)
}

// TestingObjectStore is a testing implementation of the ObjectStore interface.
type TestingObjectStore struct{}

// Get returns an io.ReadCloser for data at path, namespaced to the
// model.
func (TestingObjectStore) Get(ctx context.Context, name string) (io.ReadCloser, int64, error) {
	return nil, 0, errors.NotFoundf(name)
}

// Put stores data from reader at path, namespaced to the model.
func (TestingObjectStore) Put(ctx context.Context, path string, r io.Reader, size int64) error {
	return nil
}

// Put stores data from reader at path, namespaced to the model.
// It also ensures the stored data has the correct hash.
func (TestingObjectStore) PutAndCheckHash(ctx context.Context, path string, r io.Reader, size int64, hash string) error {
	return nil
}

// Remove removes data at path, namespaced to the model.
func (TestingObjectStore) Remove(ctx context.Context, path string) error {
	return nil
}
