// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/storage (interfaces: ModelStorageRegistryGetter)
//
// Generated by this command:
//
//	mockgen -typed -package migration_test -destination storage_mock_test.go github.com/juju/juju/core/storage ModelStorageRegistryGetter
//

// Package migration_test is a generated GoMock package.
package migration_test

import (
	context "context"
	reflect "reflect"

	storage "github.com/juju/juju/internal/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockModelStorageRegistryGetter is a mock of ModelStorageRegistryGetter interface.
type MockModelStorageRegistryGetter struct {
	ctrl     *gomock.Controller
	recorder *MockModelStorageRegistryGetterMockRecorder
}

// MockModelStorageRegistryGetterMockRecorder is the mock recorder for MockModelStorageRegistryGetter.
type MockModelStorageRegistryGetterMockRecorder struct {
	mock *MockModelStorageRegistryGetter
}

// NewMockModelStorageRegistryGetter creates a new mock instance.
func NewMockModelStorageRegistryGetter(ctrl *gomock.Controller) *MockModelStorageRegistryGetter {
	mock := &MockModelStorageRegistryGetter{ctrl: ctrl}
	mock.recorder = &MockModelStorageRegistryGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelStorageRegistryGetter) EXPECT() *MockModelStorageRegistryGetterMockRecorder {
	return m.recorder
}

// GetStorageRegistry mocks base method.
func (m *MockModelStorageRegistryGetter) GetStorageRegistry(arg0 context.Context) (storage.ProviderRegistry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageRegistry", arg0)
	ret0, _ := ret[0].(storage.ProviderRegistry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageRegistry indicates an expected call of GetStorageRegistry.
func (mr *MockModelStorageRegistryGetterMockRecorder) GetStorageRegistry(arg0 any) *MockModelStorageRegistryGetterGetStorageRegistryCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageRegistry", reflect.TypeOf((*MockModelStorageRegistryGetter)(nil).GetStorageRegistry), arg0)
	return &MockModelStorageRegistryGetterGetStorageRegistryCall{Call: call}
}

// MockModelStorageRegistryGetterGetStorageRegistryCall wrap *gomock.Call
type MockModelStorageRegistryGetterGetStorageRegistryCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelStorageRegistryGetterGetStorageRegistryCall) Return(arg0 storage.ProviderRegistry, arg1 error) *MockModelStorageRegistryGetterGetStorageRegistryCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelStorageRegistryGetterGetStorageRegistryCall) Do(f func(context.Context) (storage.ProviderRegistry, error)) *MockModelStorageRegistryGetterGetStorageRegistryCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelStorageRegistryGetterGetStorageRegistryCall) DoAndReturn(f func(context.Context) (storage.ProviderRegistry, error)) *MockModelStorageRegistryGetterGetStorageRegistryCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
