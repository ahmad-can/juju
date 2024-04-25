// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/secretbackends (interfaces: SecretBackendService)
//
// Generated by this command:
//
//	mockgen -package secretbackends -destination mock_service.go github.com/juju/juju/apiserver/facades/client/secretbackends SecretBackendService
//

// Package secretbackends is a generated GoMock package.
package secretbackends

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	service "github.com/juju/juju/domain/secretbackend/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretBackendService is a mock of SecretBackendService interface.
type MockSecretBackendService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretBackendServiceMockRecorder
}

// MockSecretBackendServiceMockRecorder is the mock recorder for MockSecretBackendService.
type MockSecretBackendServiceMockRecorder struct {
	mock *MockSecretBackendService
}

// NewMockSecretBackendService creates a new mock instance.
func NewMockSecretBackendService(ctrl *gomock.Controller) *MockSecretBackendService {
	mock := &MockSecretBackendService{ctrl: ctrl}
	mock.recorder = &MockSecretBackendServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretBackendService) EXPECT() *MockSecretBackendServiceMockRecorder {
	return m.recorder
}

// BackendSummaryInfo mocks base method.
func (m *MockSecretBackendService) BackendSummaryInfo(arg0 context.Context, arg1 bool, arg2 ...string) ([]*service.SecretBackendInfo, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BackendSummaryInfo", varargs...)
	ret0, _ := ret[0].([]*service.SecretBackendInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BackendSummaryInfo indicates an expected call of BackendSummaryInfo.
func (mr *MockSecretBackendServiceMockRecorder) BackendSummaryInfo(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackendSummaryInfo", reflect.TypeOf((*MockSecretBackendService)(nil).BackendSummaryInfo), varargs...)
}

// CreateSecretBackend mocks base method.
func (m *MockSecretBackendService) CreateSecretBackend(arg0 context.Context, arg1 secrets.SecretBackend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecretBackend indicates an expected call of CreateSecretBackend.
func (mr *MockSecretBackendServiceMockRecorder) CreateSecretBackend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecretBackend", reflect.TypeOf((*MockSecretBackendService)(nil).CreateSecretBackend), arg0, arg1)
}

// DeleteSecretBackend mocks base method.
func (m *MockSecretBackendService) DeleteSecretBackend(arg0 context.Context, arg1 service.DeleteSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecretBackend indicates an expected call of DeleteSecretBackend.
func (mr *MockSecretBackendServiceMockRecorder) DeleteSecretBackend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecretBackend", reflect.TypeOf((*MockSecretBackendService)(nil).DeleteSecretBackend), arg0, arg1)
}

// GetSecretBackendByName mocks base method.
func (m *MockSecretBackendService) GetSecretBackendByName(arg0 context.Context, arg1 string) (*secrets.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackendByName", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendByName indicates an expected call of GetSecretBackendByName.
func (mr *MockSecretBackendServiceMockRecorder) GetSecretBackendByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendByName", reflect.TypeOf((*MockSecretBackendService)(nil).GetSecretBackendByName), arg0, arg1)
}

// UpdateSecretBackend mocks base method.
func (m *MockSecretBackendService) UpdateSecretBackend(arg0 context.Context, arg1 service.UpdateSecretBackendParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecretBackend indicates an expected call of UpdateSecretBackend.
func (mr *MockSecretBackendServiceMockRecorder) UpdateSecretBackend(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretBackend", reflect.TypeOf((*MockSecretBackendService)(nil).UpdateSecretBackend), arg0, arg1)
}
