// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/upgradesteps (interfaces: UpgradeStepsState,Machine,Unit)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	controller "github.com/juju/juju/controller"
	instance "github.com/juju/juju/core/instance"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
)

// MockUpgradeStepsState is a mock of UpgradeStepsState interface.
type MockUpgradeStepsState struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeStepsStateMockRecorder
}

// MockUpgradeStepsStateMockRecorder is the mock recorder for MockUpgradeStepsState.
type MockUpgradeStepsStateMockRecorder struct {
	mock *MockUpgradeStepsState
}

// NewMockUpgradeStepsState creates a new mock instance.
func NewMockUpgradeStepsState(ctrl *gomock.Controller) *MockUpgradeStepsState {
	mock := &MockUpgradeStepsState{ctrl: ctrl}
	mock.recorder = &MockUpgradeStepsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeStepsState) EXPECT() *MockUpgradeStepsStateMockRecorder {
	return m.recorder
}

// ApplyOperation mocks base method.
func (m *MockUpgradeStepsState) ApplyOperation(arg0 state.ModelOperation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyOperation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyOperation indicates an expected call of ApplyOperation.
func (mr *MockUpgradeStepsStateMockRecorder) ApplyOperation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyOperation", reflect.TypeOf((*MockUpgradeStepsState)(nil).ApplyOperation), arg0)
}

// ControllerConfig mocks base method.
func (m *MockUpgradeStepsState) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockUpgradeStepsStateMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockUpgradeStepsState)(nil).ControllerConfig))
}

// FindEntity mocks base method.
func (m *MockUpgradeStepsState) FindEntity(arg0 names.Tag) (state.Entity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEntity", arg0)
	ret0, _ := ret[0].(state.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEntity indicates an expected call of FindEntity.
func (mr *MockUpgradeStepsStateMockRecorder) FindEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEntity", reflect.TypeOf((*MockUpgradeStepsState)(nil).FindEntity), arg0)
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

// ContainerType mocks base method.
func (m *MockMachine) ContainerType() instance.ContainerType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType.
func (mr *MockMachineMockRecorder) ContainerType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockMachine)(nil).ContainerType))
}

// ModificationStatus mocks base method.
func (m *MockMachine) ModificationStatus() (status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModificationStatus")
	ret0, _ := ret[0].(status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModificationStatus indicates an expected call of ModificationStatus.
func (mr *MockMachineMockRecorder) ModificationStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModificationStatus", reflect.TypeOf((*MockMachine)(nil).ModificationStatus))
}

// SetModificationStatus mocks base method.
func (m *MockMachine) SetModificationStatus(arg0 status.StatusInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModificationStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModificationStatus indicates an expected call of SetModificationStatus.
func (mr *MockMachineMockRecorder) SetModificationStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModificationStatus", reflect.TypeOf((*MockMachine)(nil).SetModificationStatus), arg0)
}

// MockUnit is a mock of Unit interface.
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit.
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance.
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// SetStateOperation mocks base method.
func (m *MockUnit) SetStateOperation(arg0 *state.UnitState, arg1 state.UnitStateSizeLimits) state.ModelOperation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStateOperation", arg0, arg1)
	ret0, _ := ret[0].(state.ModelOperation)
	return ret0
}

// SetStateOperation indicates an expected call of SetStateOperation.
func (mr *MockUnitMockRecorder) SetStateOperation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateOperation", reflect.TypeOf((*MockUnit)(nil).SetStateOperation), arg0, arg1)
}
