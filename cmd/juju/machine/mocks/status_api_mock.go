// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/machine (interfaces: StatusAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	params "github.com/juju/juju/rpc/params"
)

// MockStatusAPI is a mock of StatusAPI interface.
type MockStatusAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStatusAPIMockRecorder
}

// MockStatusAPIMockRecorder is the mock recorder for MockStatusAPI.
type MockStatusAPIMockRecorder struct {
	mock *MockStatusAPI
}

// NewMockStatusAPI creates a new mock instance.
func NewMockStatusAPI(ctrl *gomock.Controller) *MockStatusAPI {
	mock := &MockStatusAPI{ctrl: ctrl}
	mock.recorder = &MockStatusAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusAPI) EXPECT() *MockStatusAPIMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockStatusAPI) Status(arg0 []string) (*params.FullStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*params.FullStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockStatusAPIMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStatusAPI)(nil).Status), arg0)
}
