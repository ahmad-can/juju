// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common/storagecommon (interfaces: VolumeAccess)
//
// Generated by this command:
//
//	mockgen -typed -package machinemanager -destination volume_access_mock_test.go github.com/juju/juju/apiserver/common/storagecommon VolumeAccess
//

// Package machinemanager is a generated GoMock package.
package machinemanager

import (
	reflect "reflect"

	state "github.com/juju/juju/state"
	names "github.com/juju/names/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockVolumeAccess is a mock of VolumeAccess interface.
type MockVolumeAccess struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeAccessMockRecorder
}

// MockVolumeAccessMockRecorder is the mock recorder for MockVolumeAccess.
type MockVolumeAccessMockRecorder struct {
	mock *MockVolumeAccess
}

// NewMockVolumeAccess creates a new mock instance.
func NewMockVolumeAccess(ctrl *gomock.Controller) *MockVolumeAccess {
	mock := &MockVolumeAccess{ctrl: ctrl}
	mock.recorder = &MockVolumeAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeAccess) EXPECT() *MockVolumeAccessMockRecorder {
	return m.recorder
}

// StorageInstanceVolume mocks base method.
func (m *MockVolumeAccess) StorageInstanceVolume(arg0 names.StorageTag) (state.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageInstanceVolume", arg0)
	ret0, _ := ret[0].(state.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageInstanceVolume indicates an expected call of StorageInstanceVolume.
func (mr *MockVolumeAccessMockRecorder) StorageInstanceVolume(arg0 any) *MockVolumeAccessStorageInstanceVolumeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageInstanceVolume", reflect.TypeOf((*MockVolumeAccess)(nil).StorageInstanceVolume), arg0)
	return &MockVolumeAccessStorageInstanceVolumeCall{Call: call}
}

// MockVolumeAccessStorageInstanceVolumeCall wrap *gomock.Call
type MockVolumeAccessStorageInstanceVolumeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockVolumeAccessStorageInstanceVolumeCall) Return(arg0 state.Volume, arg1 error) *MockVolumeAccessStorageInstanceVolumeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockVolumeAccessStorageInstanceVolumeCall) Do(f func(names.StorageTag) (state.Volume, error)) *MockVolumeAccessStorageInstanceVolumeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockVolumeAccessStorageInstanceVolumeCall) DoAndReturn(f func(names.StorageTag) (state.Volume, error)) *MockVolumeAccessStorageInstanceVolumeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VolumeAttachment mocks base method.
func (m *MockVolumeAccess) VolumeAttachment(arg0 names.Tag, arg1 names.VolumeTag) (state.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeAttachment", arg0, arg1)
	ret0, _ := ret[0].(state.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeAttachment indicates an expected call of VolumeAttachment.
func (mr *MockVolumeAccessMockRecorder) VolumeAttachment(arg0, arg1 any) *MockVolumeAccessVolumeAttachmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeAttachment", reflect.TypeOf((*MockVolumeAccess)(nil).VolumeAttachment), arg0, arg1)
	return &MockVolumeAccessVolumeAttachmentCall{Call: call}
}

// MockVolumeAccessVolumeAttachmentCall wrap *gomock.Call
type MockVolumeAccessVolumeAttachmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockVolumeAccessVolumeAttachmentCall) Return(arg0 state.VolumeAttachment, arg1 error) *MockVolumeAccessVolumeAttachmentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockVolumeAccessVolumeAttachmentCall) Do(f func(names.Tag, names.VolumeTag) (state.VolumeAttachment, error)) *MockVolumeAccessVolumeAttachmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockVolumeAccessVolumeAttachmentCall) DoAndReturn(f func(names.Tag, names.VolumeTag) (state.VolumeAttachment, error)) *MockVolumeAccessVolumeAttachmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VolumeAttachmentPlan mocks base method.
func (m *MockVolumeAccess) VolumeAttachmentPlan(arg0 names.Tag, arg1 names.VolumeTag) (state.VolumeAttachmentPlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeAttachmentPlan", arg0, arg1)
	ret0, _ := ret[0].(state.VolumeAttachmentPlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeAttachmentPlan indicates an expected call of VolumeAttachmentPlan.
func (mr *MockVolumeAccessMockRecorder) VolumeAttachmentPlan(arg0, arg1 any) *MockVolumeAccessVolumeAttachmentPlanCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeAttachmentPlan", reflect.TypeOf((*MockVolumeAccess)(nil).VolumeAttachmentPlan), arg0, arg1)
	return &MockVolumeAccessVolumeAttachmentPlanCall{Call: call}
}

// MockVolumeAccessVolumeAttachmentPlanCall wrap *gomock.Call
type MockVolumeAccessVolumeAttachmentPlanCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockVolumeAccessVolumeAttachmentPlanCall) Return(arg0 state.VolumeAttachmentPlan, arg1 error) *MockVolumeAccessVolumeAttachmentPlanCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockVolumeAccessVolumeAttachmentPlanCall) Do(f func(names.Tag, names.VolumeTag) (state.VolumeAttachmentPlan, error)) *MockVolumeAccessVolumeAttachmentPlanCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockVolumeAccessVolumeAttachmentPlanCall) DoAndReturn(f func(names.Tag, names.VolumeTag) (state.VolumeAttachmentPlan, error)) *MockVolumeAccessVolumeAttachmentPlanCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
