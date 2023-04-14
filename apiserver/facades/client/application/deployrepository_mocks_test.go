// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/application (interfaces: Bindings,DeployFromRepositoryState,DeployFromRepositoryValidator,Model,Machine)

// Package application is a generated GoMock package.
package application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v10 "github.com/juju/charm/v10"
	resource "github.com/juju/charm/v10/resource"
	services "github.com/juju/juju/apiserver/facades/client/charms/services"
	cloud "github.com/juju/juju/cloud"
	controller "github.com/juju/juju/controller"
	constraints "github.com/juju/juju/core/constraints"
	instance "github.com/juju/juju/core/instance"
	network "github.com/juju/juju/core/network"
	config "github.com/juju/juju/environs/config"
	params "github.com/juju/juju/rpc/params"
	state "github.com/juju/juju/state"
	v4 "github.com/juju/names/v4"
	v2 "github.com/juju/version/v2"
)

// MockBindings is a mock of Bindings interface.
type MockBindings struct {
	ctrl     *gomock.Controller
	recorder *MockBindingsMockRecorder
}

// MockBindingsMockRecorder is the mock recorder for MockBindings.
type MockBindingsMockRecorder struct {
	mock *MockBindings
}

// NewMockBindings creates a new mock instance.
func NewMockBindings(ctrl *gomock.Controller) *MockBindings {
	mock := &MockBindings{ctrl: ctrl}
	mock.recorder = &MockBindingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBindings) EXPECT() *MockBindingsMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockBindings) Map() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockBindingsMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockBindings)(nil).Map))
}

// MapWithSpaceNames mocks base method.
func (m *MockBindings) MapWithSpaceNames(arg0 network.SpaceInfos) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MapWithSpaceNames", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapWithSpaceNames indicates an expected call of MapWithSpaceNames.
func (mr *MockBindingsMockRecorder) MapWithSpaceNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapWithSpaceNames", reflect.TypeOf((*MockBindings)(nil).MapWithSpaceNames), arg0)
}

// MockDeployFromRepositoryState is a mock of DeployFromRepositoryState interface.
type MockDeployFromRepositoryState struct {
	ctrl     *gomock.Controller
	recorder *MockDeployFromRepositoryStateMockRecorder
}

// MockDeployFromRepositoryStateMockRecorder is the mock recorder for MockDeployFromRepositoryState.
type MockDeployFromRepositoryStateMockRecorder struct {
	mock *MockDeployFromRepositoryState
}

// NewMockDeployFromRepositoryState creates a new mock instance.
func NewMockDeployFromRepositoryState(ctrl *gomock.Controller) *MockDeployFromRepositoryState {
	mock := &MockDeployFromRepositoryState{ctrl: ctrl}
	mock.recorder = &MockDeployFromRepositoryStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployFromRepositoryState) EXPECT() *MockDeployFromRepositoryStateMockRecorder {
	return m.recorder
}

// AddApplication mocks base method.
func (m *MockDeployFromRepositoryState) AddApplication(arg0 state.AddApplicationArgs) (Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddApplication", arg0)
	ret0, _ := ret[0].(Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddApplication indicates an expected call of AddApplication.
func (mr *MockDeployFromRepositoryStateMockRecorder) AddApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApplication", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).AddApplication), arg0)
}

// AddCharmMetadata mocks base method.
func (m *MockDeployFromRepositoryState) AddCharmMetadata(arg0 state.CharmInfo) (Charm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharmMetadata", arg0)
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharmMetadata indicates an expected call of AddCharmMetadata.
func (mr *MockDeployFromRepositoryStateMockRecorder) AddCharmMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharmMetadata", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).AddCharmMetadata), arg0)
}

// AddPendingResource mocks base method.
func (m *MockDeployFromRepositoryState) AddPendingResource(arg0 string, arg1 resource.Resource) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPendingResource", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPendingResource indicates an expected call of AddPendingResource.
func (mr *MockDeployFromRepositoryStateMockRecorder) AddPendingResource(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPendingResource", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).AddPendingResource), arg0, arg1)
}

// AllSpaceInfos mocks base method.
func (m *MockDeployFromRepositoryState) AllSpaceInfos() (network.SpaceInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllSpaceInfos")
	ret0, _ := ret[0].(network.SpaceInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllSpaceInfos indicates an expected call of AllSpaceInfos.
func (mr *MockDeployFromRepositoryStateMockRecorder) AllSpaceInfos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllSpaceInfos", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).AllSpaceInfos))
}

// ControllerConfig mocks base method.
func (m *MockDeployFromRepositoryState) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockDeployFromRepositoryStateMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).ControllerConfig))
}

// DefaultEndpointBindingSpace mocks base method.
func (m *MockDeployFromRepositoryState) DefaultEndpointBindingSpace() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultEndpointBindingSpace")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultEndpointBindingSpace indicates an expected call of DefaultEndpointBindingSpace.
func (mr *MockDeployFromRepositoryStateMockRecorder) DefaultEndpointBindingSpace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultEndpointBindingSpace", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).DefaultEndpointBindingSpace))
}

// Machine mocks base method.
func (m *MockDeployFromRepositoryState) Machine(arg0 string) (Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockDeployFromRepositoryStateMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).Machine), arg0)
}

// ModelConstraints mocks base method.
func (m *MockDeployFromRepositoryState) ModelConstraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConstraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConstraints indicates an expected call of ModelConstraints.
func (mr *MockDeployFromRepositoryStateMockRecorder) ModelConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConstraints", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).ModelConstraints))
}

// ModelUUID mocks base method.
func (m *MockDeployFromRepositoryState) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockDeployFromRepositoryStateMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).ModelUUID))
}

// PrepareCharmUpload mocks base method.
func (m *MockDeployFromRepositoryState) PrepareCharmUpload(arg0 *v10.URL) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockDeployFromRepositoryStateMockRecorder) PrepareCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).PrepareCharmUpload), arg0)
}

// RemovePendingResources mocks base method.
func (m *MockDeployFromRepositoryState) RemovePendingResources(arg0 string, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePendingResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePendingResources indicates an expected call of RemovePendingResources.
func (mr *MockDeployFromRepositoryStateMockRecorder) RemovePendingResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePendingResources", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).RemovePendingResources), arg0, arg1)
}

// Space mocks base method.
func (m *MockDeployFromRepositoryState) Space(arg0 string) (*state.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Space", arg0)
	ret0, _ := ret[0].(*state.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Space indicates an expected call of Space.
func (mr *MockDeployFromRepositoryStateMockRecorder) Space(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Space", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).Space), arg0)
}

// UpdateUploadedCharm mocks base method.
func (m *MockDeployFromRepositoryState) UpdateUploadedCharm(arg0 state.CharmInfo) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockDeployFromRepositoryStateMockRecorder) UpdateUploadedCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockDeployFromRepositoryState)(nil).UpdateUploadedCharm), arg0)
}

// MockDeployFromRepositoryValidator is a mock of DeployFromRepositoryValidator interface.
type MockDeployFromRepositoryValidator struct {
	ctrl     *gomock.Controller
	recorder *MockDeployFromRepositoryValidatorMockRecorder
}

// MockDeployFromRepositoryValidatorMockRecorder is the mock recorder for MockDeployFromRepositoryValidator.
type MockDeployFromRepositoryValidatorMockRecorder struct {
	mock *MockDeployFromRepositoryValidator
}

// NewMockDeployFromRepositoryValidator creates a new mock instance.
func NewMockDeployFromRepositoryValidator(ctrl *gomock.Controller) *MockDeployFromRepositoryValidator {
	mock := &MockDeployFromRepositoryValidator{ctrl: ctrl}
	mock.recorder = &MockDeployFromRepositoryValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployFromRepositoryValidator) EXPECT() *MockDeployFromRepositoryValidatorMockRecorder {
	return m.recorder
}

// ValidateArg mocks base method.
func (m *MockDeployFromRepositoryValidator) ValidateArg(arg0 params.DeployFromRepositoryArg) (deployTemplate, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateArg", arg0)
	ret0, _ := ret[0].(deployTemplate)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// ValidateArg indicates an expected call of ValidateArg.
func (mr *MockDeployFromRepositoryValidatorMockRecorder) ValidateArg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateArg", reflect.TypeOf((*MockDeployFromRepositoryValidator)(nil).ValidateArg), arg0)
}

// MockModel is a mock of Model interface.
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel.
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance.
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// AgentVersion mocks base method.
func (m *MockModel) AgentVersion() (v2.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion")
	ret0, _ := ret[0].(v2.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion.
func (mr *MockModelMockRecorder) AgentVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockModel)(nil).AgentVersion))
}

// Cloud mocks base method.
func (m *MockModel) Cloud() (cloud.Cloud, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cloud")
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cloud indicates an expected call of Cloud.
func (mr *MockModelMockRecorder) Cloud() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cloud", reflect.TypeOf((*MockModel)(nil).Cloud))
}

// CloudCredential mocks base method.
func (m *MockModel) CloudCredential() (state.Credential, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudCredential")
	ret0, _ := ret[0].(state.Credential)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CloudCredential indicates an expected call of CloudCredential.
func (mr *MockModelMockRecorder) CloudCredential() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudCredential", reflect.TypeOf((*MockModel)(nil).CloudCredential))
}

// CloudName mocks base method.
func (m *MockModel) CloudName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudName indicates an expected call of CloudName.
func (mr *MockModelMockRecorder) CloudName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudName", reflect.TypeOf((*MockModel)(nil).CloudName))
}

// CloudRegion mocks base method.
func (m *MockModel) CloudRegion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudRegion")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudRegion indicates an expected call of CloudRegion.
func (mr *MockModelMockRecorder) CloudRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudRegion", reflect.TypeOf((*MockModel)(nil).CloudRegion))
}

// Config mocks base method.
func (m *MockModel) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModel)(nil).Config))
}

// ControllerUUID mocks base method.
func (m *MockModel) ControllerUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ControllerUUID indicates an expected call of ControllerUUID.
func (mr *MockModelMockRecorder) ControllerUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerUUID", reflect.TypeOf((*MockModel)(nil).ControllerUUID))
}

// ModelConfig mocks base method.
func (m *MockModel) ModelConfig() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelMockRecorder) ModelConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModel)(nil).ModelConfig))
}

// ModelTag mocks base method.
func (m *MockModel) ModelTag() v4.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(v4.ModelTag)
	return ret0
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockModelMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockModel)(nil).ModelTag))
}

// Name mocks base method.
func (m *MockModel) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockModelMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModel)(nil).Name))
}

// OpenedPortRangesForMachine mocks base method.
func (m *MockModel) OpenedPortRangesForMachine(arg0 string) (state.MachinePortRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenedPortRangesForMachine", arg0)
	ret0, _ := ret[0].(state.MachinePortRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenedPortRangesForMachine indicates an expected call of OpenedPortRangesForMachine.
func (mr *MockModelMockRecorder) OpenedPortRangesForMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenedPortRangesForMachine", reflect.TypeOf((*MockModel)(nil).OpenedPortRangesForMachine), arg0)
}

// Owner mocks base method.
func (m *MockModel) Owner() v4.UserTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(v4.UserTag)
	return ret0
}

// Owner indicates an expected call of Owner.
func (mr *MockModelMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockModel)(nil).Owner))
}

// Tag mocks base method.
func (m *MockModel) Tag() v4.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(v4.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockModelMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockModel)(nil).Tag))
}

// Type mocks base method.
func (m *MockModel) Type() state.ModelType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.ModelType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockModelMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockModel)(nil).Type))
}

// UUID mocks base method.
func (m *MockModel) UUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UUID indicates an expected call of UUID.
func (mr *MockModelMockRecorder) UUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UUID", reflect.TypeOf((*MockModel)(nil).UUID))
}

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// Base mocks base method.
func (m *MockMachine) Base() state.Base {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Base")
	ret0, _ := ret[0].(state.Base)
	return ret0
}

// Base indicates an expected call of Base.
func (mr *MockMachineMockRecorder) Base() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Base", reflect.TypeOf((*MockMachine)(nil).Base))
}

// HardwareCharacteristics mocks base method.
func (m *MockMachine) HardwareCharacteristics() (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics")
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics.
func (mr *MockMachineMockRecorder) HardwareCharacteristics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockMachine)(nil).HardwareCharacteristics))
}

// IsLockedForSeriesUpgrade mocks base method.
func (m *MockMachine) IsLockedForSeriesUpgrade() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLockedForSeriesUpgrade")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLockedForSeriesUpgrade indicates an expected call of IsLockedForSeriesUpgrade.
func (mr *MockMachineMockRecorder) IsLockedForSeriesUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLockedForSeriesUpgrade", reflect.TypeOf((*MockMachine)(nil).IsLockedForSeriesUpgrade))
}

// IsParentLockedForSeriesUpgrade mocks base method.
func (m *MockMachine) IsParentLockedForSeriesUpgrade() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsParentLockedForSeriesUpgrade")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsParentLockedForSeriesUpgrade indicates an expected call of IsParentLockedForSeriesUpgrade.
func (mr *MockMachineMockRecorder) IsParentLockedForSeriesUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsParentLockedForSeriesUpgrade", reflect.TypeOf((*MockMachine)(nil).IsParentLockedForSeriesUpgrade))
}

// PublicAddress mocks base method.
func (m *MockMachine) PublicAddress() (network.SpaceAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublicAddress")
	ret0, _ := ret[0].(network.SpaceAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublicAddress indicates an expected call of PublicAddress.
func (mr *MockMachineMockRecorder) PublicAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublicAddress", reflect.TypeOf((*MockMachine)(nil).PublicAddress))
}
