// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api/base (interfaces: APICaller)
//
// Generated by this command:
//
//	mockgen -typed -package computeprovisioner_test -destination base_mock_test.go github.com/juju/juju/api/base APICaller
//

// Package computeprovisioner_test is a generated GoMock package.
package computeprovisioner_test

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	base "github.com/juju/juju/api/base"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
)

// MockAPICaller is a mock of APICaller interface.
type MockAPICaller struct {
	ctrl     *gomock.Controller
	recorder *MockAPICallerMockRecorder
}

// MockAPICallerMockRecorder is the mock recorder for MockAPICaller.
type MockAPICallerMockRecorder struct {
	mock *MockAPICaller
}

// NewMockAPICaller creates a new mock instance.
func NewMockAPICaller(ctrl *gomock.Controller) *MockAPICaller {
	mock := &MockAPICaller{ctrl: ctrl}
	mock.recorder = &MockAPICallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPICaller) EXPECT() *MockAPICallerMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockAPICaller) APICall(arg0 context.Context, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockAPICallerMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *MockAPICallerAPICallCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockAPICaller)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	return &MockAPICallerAPICallCall{Call: call}
}

// MockAPICallerAPICallCall wrap *gomock.Call
type MockAPICallerAPICallCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerAPICallCall) Return(arg0 error) *MockAPICallerAPICallCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerAPICallCall) Do(f func(context.Context, string, int, string, string, any, any) error) *MockAPICallerAPICallCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerAPICallCall) DoAndReturn(f func(context.Context, string, int, string, string, any, any) error) *MockAPICallerAPICallCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BakeryClient mocks base method.
func (m *MockAPICaller) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockAPICallerMockRecorder) BakeryClient() *MockAPICallerBakeryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockAPICaller)(nil).BakeryClient))
	return &MockAPICallerBakeryClientCall{Call: call}
}

// MockAPICallerBakeryClientCall wrap *gomock.Call
type MockAPICallerBakeryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerBakeryClientCall) Return(arg0 base.MacaroonDischarger) *MockAPICallerBakeryClientCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerBakeryClientCall) Do(f func() base.MacaroonDischarger) *MockAPICallerBakeryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerBakeryClientCall) DoAndReturn(f func() base.MacaroonDischarger) *MockAPICallerBakeryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// BestFacadeVersion mocks base method.
func (m *MockAPICaller) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockAPICallerMockRecorder) BestFacadeVersion(arg0 any) *MockAPICallerBestFacadeVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockAPICaller)(nil).BestFacadeVersion), arg0)
	return &MockAPICallerBestFacadeVersionCall{Call: call}
}

// MockAPICallerBestFacadeVersionCall wrap *gomock.Call
type MockAPICallerBestFacadeVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerBestFacadeVersionCall) Return(arg0 int) *MockAPICallerBestFacadeVersionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerBestFacadeVersionCall) Do(f func(string) int) *MockAPICallerBestFacadeVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerBestFacadeVersionCall) DoAndReturn(f func(string) int) *MockAPICallerBestFacadeVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectControllerStream mocks base method.
func (m *MockAPICaller) ConnectControllerStream(arg0 context.Context, arg1 string, arg2 url.Values, arg3 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockAPICallerMockRecorder) ConnectControllerStream(arg0, arg1, arg2, arg3 any) *MockAPICallerConnectControllerStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockAPICaller)(nil).ConnectControllerStream), arg0, arg1, arg2, arg3)
	return &MockAPICallerConnectControllerStreamCall{Call: call}
}

// MockAPICallerConnectControllerStreamCall wrap *gomock.Call
type MockAPICallerConnectControllerStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerConnectControllerStreamCall) Return(arg0 base.Stream, arg1 error) *MockAPICallerConnectControllerStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerConnectControllerStreamCall) Do(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockAPICallerConnectControllerStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerConnectControllerStreamCall) DoAndReturn(f func(context.Context, string, url.Values, http.Header) (base.Stream, error)) *MockAPICallerConnectControllerStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectStream mocks base method.
func (m *MockAPICaller) ConnectStream(arg0 context.Context, arg1 string, arg2 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockAPICallerMockRecorder) ConnectStream(arg0, arg1, arg2 any) *MockAPICallerConnectStreamCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockAPICaller)(nil).ConnectStream), arg0, arg1, arg2)
	return &MockAPICallerConnectStreamCall{Call: call}
}

// MockAPICallerConnectStreamCall wrap *gomock.Call
type MockAPICallerConnectStreamCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerConnectStreamCall) Return(arg0 base.Stream, arg1 error) *MockAPICallerConnectStreamCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerConnectStreamCall) Do(f func(context.Context, string, url.Values) (base.Stream, error)) *MockAPICallerConnectStreamCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerConnectStreamCall) DoAndReturn(f func(context.Context, string, url.Values) (base.Stream, error)) *MockAPICallerConnectStreamCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HTTPClient mocks base method.
func (m *MockAPICaller) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockAPICallerMockRecorder) HTTPClient() *MockAPICallerHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockAPICaller)(nil).HTTPClient))
	return &MockAPICallerHTTPClientCall{Call: call}
}

// MockAPICallerHTTPClientCall wrap *gomock.Call
type MockAPICallerHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockAPICallerHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockAPICallerHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockAPICallerHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelTag mocks base method.
func (m *MockAPICaller) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockAPICallerMockRecorder) ModelTag() *MockAPICallerModelTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockAPICaller)(nil).ModelTag))
	return &MockAPICallerModelTagCall{Call: call}
}

// MockAPICallerModelTagCall wrap *gomock.Call
type MockAPICallerModelTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerModelTagCall) Return(arg0 names.ModelTag, arg1 bool) *MockAPICallerModelTagCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerModelTagCall) Do(f func() (names.ModelTag, bool)) *MockAPICallerModelTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerModelTagCall) DoAndReturn(f func() (names.ModelTag, bool)) *MockAPICallerModelTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RootHTTPClient mocks base method.
func (m *MockAPICaller) RootHTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHTTPClient indicates an expected call of RootHTTPClient.
func (mr *MockAPICallerMockRecorder) RootHTTPClient() *MockAPICallerRootHTTPClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHTTPClient", reflect.TypeOf((*MockAPICaller)(nil).RootHTTPClient))
	return &MockAPICallerRootHTTPClientCall{Call: call}
}

// MockAPICallerRootHTTPClientCall wrap *gomock.Call
type MockAPICallerRootHTTPClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAPICallerRootHTTPClientCall) Return(arg0 *httprequest.Client, arg1 error) *MockAPICallerRootHTTPClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAPICallerRootHTTPClientCall) Do(f func() (*httprequest.Client, error)) *MockAPICallerRootHTTPClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAPICallerRootHTTPClientCall) DoAndReturn(f func() (*httprequest.Client, error)) *MockAPICallerRootHTTPClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
