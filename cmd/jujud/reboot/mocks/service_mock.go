// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/jujud/reboot (interfaces: AgentConfig,Manager,Model,RebootWaiter,Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	reboot "github.com/juju/juju/cmd/jujud/reboot"
	container "github.com/juju/juju/container"
	instance "github.com/juju/juju/core/instance"
	instances "github.com/juju/juju/environs/instances"
	params "github.com/juju/juju/rpc/params"
	common "github.com/juju/juju/service/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAgentConfig is a mock of AgentConfig interface.
type MockAgentConfig struct {
	ctrl     *gomock.Controller
	recorder *MockAgentConfigMockRecorder
}

// MockAgentConfigMockRecorder is the mock recorder for MockAgentConfig.
type MockAgentConfigMockRecorder struct {
	mock *MockAgentConfig
}

// NewMockAgentConfig creates a new mock instance.
func NewMockAgentConfig(ctrl *gomock.Controller) *MockAgentConfig {
	mock := &MockAgentConfig{ctrl: ctrl}
	mock.recorder = &MockAgentConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentConfig) EXPECT() *MockAgentConfigMockRecorder {
	return m.recorder
}

// Model mocks base method.
func (m *MockAgentConfig) Model() reboot.Model {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Model")
	ret0, _ := ret[0].(reboot.Model)
	return ret0
}

// Model indicates an expected call of Model.
func (mr *MockAgentConfigMockRecorder) Model() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockAgentConfig)(nil).Model))
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// IsInitialized mocks base method.
func (m *MockManager) IsInitialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockManagerMockRecorder) IsInitialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockManager)(nil).IsInitialized))
}

// ListContainers mocks base method.
func (m *MockManager) ListContainers() ([]instances.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers")
	ret0, _ := ret[0].([]instances.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockManagerMockRecorder) ListContainers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockManager)(nil).ListContainers))
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

// Id mocks base method.
func (m *MockModel) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockModelMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockModel)(nil).Id))
}

// MockRebootWaiter is a mock of RebootWaiter interface.
type MockRebootWaiter struct {
	ctrl     *gomock.Controller
	recorder *MockRebootWaiterMockRecorder
}

// MockRebootWaiterMockRecorder is the mock recorder for MockRebootWaiter.
type MockRebootWaiterMockRecorder struct {
	mock *MockRebootWaiter
}

// NewMockRebootWaiter creates a new mock instance.
func NewMockRebootWaiter(ctrl *gomock.Controller) *MockRebootWaiter {
	mock := &MockRebootWaiter{ctrl: ctrl}
	mock.recorder = &MockRebootWaiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRebootWaiter) EXPECT() *MockRebootWaiterMockRecorder {
	return m.recorder
}

// HostSeries mocks base method.
func (m *MockRebootWaiter) HostSeries() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HostSeries")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HostSeries indicates an expected call of HostSeries.
func (mr *MockRebootWaiterMockRecorder) HostSeries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostSeries", reflect.TypeOf((*MockRebootWaiter)(nil).HostSeries))
}

// ListServices mocks base method.
func (m *MockRebootWaiter) ListServices() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockRebootWaiterMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockRebootWaiter)(nil).ListServices))
}

// NewContainerManager mocks base method.
func (m *MockRebootWaiter) NewContainerManager(arg0 instance.ContainerType, arg1 container.ManagerConfig) (reboot.Manager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewContainerManager", arg0, arg1)
	ret0, _ := ret[0].(reboot.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewContainerManager indicates an expected call of NewContainerManager.
func (mr *MockRebootWaiterMockRecorder) NewContainerManager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewContainerManager", reflect.TypeOf((*MockRebootWaiter)(nil).NewContainerManager), arg0, arg1)
}

// NewService mocks base method.
func (m *MockRebootWaiter) NewService(arg0 string, arg1 common.Conf, arg2 string) (reboot.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewService", arg0, arg1, arg2)
	ret0, _ := ret[0].(reboot.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewService indicates an expected call of NewService.
func (mr *MockRebootWaiterMockRecorder) NewService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewService", reflect.TypeOf((*MockRebootWaiter)(nil).NewService), arg0, arg1, arg2)
}

// ScheduleAction mocks base method.
func (m *MockRebootWaiter) ScheduleAction(arg0 params.RebootAction, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleAction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleAction indicates an expected call of ScheduleAction.
func (mr *MockRebootWaiterMockRecorder) ScheduleAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleAction", reflect.TypeOf((*MockRebootWaiter)(nil).ScheduleAction), arg0, arg1)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Stop mocks base method.
func (m *MockService) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockServiceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockService)(nil).Stop))
}
