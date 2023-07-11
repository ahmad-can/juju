// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/migrations (interfaces: MigrationRemoteEntity,RemoteEntitiesSource,RemoteEntitiesModel)

// Package migrations is a generated GoMock package.
package migrations

import (
	reflect "reflect"

	v3 "github.com/juju/description/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockMigrationRemoteEntity is a mock of MigrationRemoteEntity interface.
type MockMigrationRemoteEntity struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationRemoteEntityMockRecorder
}

// MockMigrationRemoteEntityMockRecorder is the mock recorder for MockMigrationRemoteEntity.
type MockMigrationRemoteEntityMockRecorder struct {
	mock *MockMigrationRemoteEntity
}

// NewMockMigrationRemoteEntity creates a new mock instance.
func NewMockMigrationRemoteEntity(ctrl *gomock.Controller) *MockMigrationRemoteEntity {
	mock := &MockMigrationRemoteEntity{ctrl: ctrl}
	mock.recorder = &MockMigrationRemoteEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrationRemoteEntity) EXPECT() *MockMigrationRemoteEntityMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockMigrationRemoteEntity) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockMigrationRemoteEntityMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockMigrationRemoteEntity)(nil).ID))
}

// Macaroon mocks base method.
func (m *MockMigrationRemoteEntity) Macaroon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Macaroon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Macaroon indicates an expected call of Macaroon.
func (mr *MockMigrationRemoteEntityMockRecorder) Macaroon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Macaroon", reflect.TypeOf((*MockMigrationRemoteEntity)(nil).Macaroon))
}

// Token mocks base method.
func (m *MockMigrationRemoteEntity) Token() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(string)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockMigrationRemoteEntityMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockMigrationRemoteEntity)(nil).Token))
}

// MockRemoteEntitiesSource is a mock of RemoteEntitiesSource interface.
type MockRemoteEntitiesSource struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteEntitiesSourceMockRecorder
}

// MockRemoteEntitiesSourceMockRecorder is the mock recorder for MockRemoteEntitiesSource.
type MockRemoteEntitiesSourceMockRecorder struct {
	mock *MockRemoteEntitiesSource
}

// NewMockRemoteEntitiesSource creates a new mock instance.
func NewMockRemoteEntitiesSource(ctrl *gomock.Controller) *MockRemoteEntitiesSource {
	mock := &MockRemoteEntitiesSource{ctrl: ctrl}
	mock.recorder = &MockRemoteEntitiesSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteEntitiesSource) EXPECT() *MockRemoteEntitiesSourceMockRecorder {
	return m.recorder
}

// AllRemoteEntities mocks base method.
func (m *MockRemoteEntitiesSource) AllRemoteEntities() ([]MigrationRemoteEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRemoteEntities")
	ret0, _ := ret[0].([]MigrationRemoteEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRemoteEntities indicates an expected call of AllRemoteEntities.
func (mr *MockRemoteEntitiesSourceMockRecorder) AllRemoteEntities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRemoteEntities", reflect.TypeOf((*MockRemoteEntitiesSource)(nil).AllRemoteEntities))
}

// MockRemoteEntitiesModel is a mock of RemoteEntitiesModel interface.
type MockRemoteEntitiesModel struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteEntitiesModelMockRecorder
}

// MockRemoteEntitiesModelMockRecorder is the mock recorder for MockRemoteEntitiesModel.
type MockRemoteEntitiesModelMockRecorder struct {
	mock *MockRemoteEntitiesModel
}

// NewMockRemoteEntitiesModel creates a new mock instance.
func NewMockRemoteEntitiesModel(ctrl *gomock.Controller) *MockRemoteEntitiesModel {
	mock := &MockRemoteEntitiesModel{ctrl: ctrl}
	mock.recorder = &MockRemoteEntitiesModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteEntitiesModel) EXPECT() *MockRemoteEntitiesModelMockRecorder {
	return m.recorder
}

// AddRemoteEntity mocks base method.
func (m *MockRemoteEntitiesModel) AddRemoteEntity(arg0 v3.RemoteEntityArgs) v3.RemoteEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteEntity", arg0)
	ret0, _ := ret[0].(v3.RemoteEntity)
	return ret0
}

// AddRemoteEntity indicates an expected call of AddRemoteEntity.
func (mr *MockRemoteEntitiesModelMockRecorder) AddRemoteEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteEntity", reflect.TypeOf((*MockRemoteEntitiesModel)(nil).AddRemoteEntity), arg0)
}
