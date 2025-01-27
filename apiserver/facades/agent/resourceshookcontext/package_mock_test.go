// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/resourceshookcontext (interfaces: ApplicationService,ResourceService)
//
// Generated by this command:
//
//	mockgen -typed -package resourceshookcontext -destination package_mock_test.go github.com/juju/juju/apiserver/facades/agent/resourceshookcontext ApplicationService,ResourceService
//

// Package resourceshookcontext is a generated GoMock package.
package resourceshookcontext

import (
	context "context"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	resource "github.com/juju/juju/core/resource"
	unit "github.com/juju/juju/core/unit"
	gomock "go.uber.org/mock/gomock"
)

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// GetApplicationIDByName mocks base method.
func (m *MockApplicationService) GetApplicationIDByName(arg0 context.Context, arg1 string) (application.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationIDByName", arg0, arg1)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationIDByName indicates an expected call of GetApplicationIDByName.
func (mr *MockApplicationServiceMockRecorder) GetApplicationIDByName(arg0, arg1 any) *MockApplicationServiceGetApplicationIDByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationIDByName", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationIDByName), arg0, arg1)
	return &MockApplicationServiceGetApplicationIDByNameCall{Call: call}
}

// MockApplicationServiceGetApplicationIDByNameCall wrap *gomock.Call
type MockApplicationServiceGetApplicationIDByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationIDByNameCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationIDByNameCall) Do(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationIDByNameCall) DoAndReturn(f func(context.Context, string) (application.ID, error)) *MockApplicationServiceGetApplicationIDByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationIDByUnitName mocks base method.
func (m *MockApplicationService) GetApplicationIDByUnitName(arg0 context.Context, arg1 unit.Name) (application.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationIDByUnitName", arg0, arg1)
	ret0, _ := ret[0].(application.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationIDByUnitName indicates an expected call of GetApplicationIDByUnitName.
func (mr *MockApplicationServiceMockRecorder) GetApplicationIDByUnitName(arg0, arg1 any) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationIDByUnitName", reflect.TypeOf((*MockApplicationService)(nil).GetApplicationIDByUnitName), arg0, arg1)
	return &MockApplicationServiceGetApplicationIDByUnitNameCall{Call: call}
}

// MockApplicationServiceGetApplicationIDByUnitNameCall wrap *gomock.Call
type MockApplicationServiceGetApplicationIDByUnitNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) Return(arg0 application.ID, arg1 error) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) Do(f func(context.Context, unit.Name) (application.ID, error)) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceGetApplicationIDByUnitNameCall) DoAndReturn(f func(context.Context, unit.Name) (application.ID, error)) *MockApplicationServiceGetApplicationIDByUnitNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockResourceService is a mock of ResourceService interface.
type MockResourceService struct {
	ctrl     *gomock.Controller
	recorder *MockResourceServiceMockRecorder
}

// MockResourceServiceMockRecorder is the mock recorder for MockResourceService.
type MockResourceServiceMockRecorder struct {
	mock *MockResourceService
}

// NewMockResourceService creates a new mock instance.
func NewMockResourceService(ctrl *gomock.Controller) *MockResourceService {
	mock := &MockResourceService{ctrl: ctrl}
	mock.recorder = &MockResourceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceService) EXPECT() *MockResourceServiceMockRecorder {
	return m.recorder
}

// GetResourcesByApplicationID mocks base method.
func (m *MockResourceService) GetResourcesByApplicationID(arg0 context.Context, arg1 application.ID) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesByApplicationID", arg0, arg1)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesByApplicationID indicates an expected call of GetResourcesByApplicationID.
func (mr *MockResourceServiceMockRecorder) GetResourcesByApplicationID(arg0, arg1 any) *MockResourceServiceGetResourcesByApplicationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesByApplicationID", reflect.TypeOf((*MockResourceService)(nil).GetResourcesByApplicationID), arg0, arg1)
	return &MockResourceServiceGetResourcesByApplicationIDCall{Call: call}
}

// MockResourceServiceGetResourcesByApplicationIDCall wrap *gomock.Call
type MockResourceServiceGetResourcesByApplicationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceServiceGetResourcesByApplicationIDCall) Return(arg0 []resource.Resource, arg1 error) *MockResourceServiceGetResourcesByApplicationIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceServiceGetResourcesByApplicationIDCall) Do(f func(context.Context, application.ID) ([]resource.Resource, error)) *MockResourceServiceGetResourcesByApplicationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceServiceGetResourcesByApplicationIDCall) DoAndReturn(f func(context.Context, application.ID) ([]resource.Resource, error)) *MockResourceServiceGetResourcesByApplicationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
