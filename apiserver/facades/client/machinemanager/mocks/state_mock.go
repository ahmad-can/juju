// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: StorageAttachment,StorageInstance,Block)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	state "github.com/juju/juju/state"
	txn "github.com/juju/mgo/v2/txn"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockStorageAttachment is a mock of StorageAttachment interface.
type MockStorageAttachment struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAttachmentMockRecorder
}

// MockStorageAttachmentMockRecorder is the mock recorder for MockStorageAttachment.
type MockStorageAttachmentMockRecorder struct {
	mock *MockStorageAttachment
}

// NewMockStorageAttachment creates a new mock instance.
func NewMockStorageAttachment(ctrl *gomock.Controller) *MockStorageAttachment {
	mock := &MockStorageAttachment{ctrl: ctrl}
	mock.recorder = &MockStorageAttachmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAttachment) EXPECT() *MockStorageAttachmentMockRecorder {
	return m.recorder
}

// Life mocks base method.
func (m *MockStorageAttachment) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockStorageAttachmentMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockStorageAttachment)(nil).Life))
}

// StorageInstance mocks base method.
func (m *MockStorageAttachment) StorageInstance() names.StorageTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageInstance")
	ret0, _ := ret[0].(names.StorageTag)
	return ret0
}

// StorageInstance indicates an expected call of StorageInstance.
func (mr *MockStorageAttachmentMockRecorder) StorageInstance() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageInstance", reflect.TypeOf((*MockStorageAttachment)(nil).StorageInstance))
}

// Unit mocks base method.
func (m *MockStorageAttachment) Unit() names.UnitTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unit")
	ret0, _ := ret[0].(names.UnitTag)
	return ret0
}

// Unit indicates an expected call of Unit.
func (mr *MockStorageAttachmentMockRecorder) Unit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unit", reflect.TypeOf((*MockStorageAttachment)(nil).Unit))
}

// MockStorageInstance is a mock of StorageInstance interface.
type MockStorageInstance struct {
	ctrl     *gomock.Controller
	recorder *MockStorageInstanceMockRecorder
}

// MockStorageInstanceMockRecorder is the mock recorder for MockStorageInstance.
type MockStorageInstanceMockRecorder struct {
	mock *MockStorageInstance
}

// NewMockStorageInstance creates a new mock instance.
func NewMockStorageInstance(ctrl *gomock.Controller) *MockStorageInstance {
	mock := &MockStorageInstance{ctrl: ctrl}
	mock.recorder = &MockStorageInstanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageInstance) EXPECT() *MockStorageInstanceMockRecorder {
	return m.recorder
}

// Kind mocks base method.
func (m *MockStorageInstance) Kind() state.StorageKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(state.StorageKind)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockStorageInstanceMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockStorageInstance)(nil).Kind))
}

// Life mocks base method.
func (m *MockStorageInstance) Life() state.Life {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life")
	ret0, _ := ret[0].(state.Life)
	return ret0
}

// Life indicates an expected call of Life.
func (mr *MockStorageInstanceMockRecorder) Life() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockStorageInstance)(nil).Life))
}

// Owner mocks base method.
func (m *MockStorageInstance) Owner() (names.Tag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Owner")
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Owner indicates an expected call of Owner.
func (mr *MockStorageInstanceMockRecorder) Owner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Owner", reflect.TypeOf((*MockStorageInstance)(nil).Owner))
}

// Pool mocks base method.
func (m *MockStorageInstance) Pool() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pool")
	ret0, _ := ret[0].(string)
	return ret0
}

// Pool indicates an expected call of Pool.
func (mr *MockStorageInstanceMockRecorder) Pool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pool", reflect.TypeOf((*MockStorageInstance)(nil).Pool))
}

// StorageName mocks base method.
func (m *MockStorageInstance) StorageName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StorageName indicates an expected call of StorageName.
func (mr *MockStorageInstanceMockRecorder) StorageName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageName", reflect.TypeOf((*MockStorageInstance)(nil).StorageName))
}

// StorageTag mocks base method.
func (m *MockStorageInstance) StorageTag() names.StorageTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageTag")
	ret0, _ := ret[0].(names.StorageTag)
	return ret0
}

// StorageTag indicates an expected call of StorageTag.
func (mr *MockStorageInstanceMockRecorder) StorageTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageTag", reflect.TypeOf((*MockStorageInstance)(nil).StorageTag))
}

// Tag mocks base method.
func (m *MockStorageInstance) Tag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockStorageInstanceMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockStorageInstance)(nil).Tag))
}

// MockBlock is a mock of Block interface.
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock.
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance.
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// Id mocks base method.
func (m *MockBlock) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockBlockMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockBlock)(nil).Id))
}

// Message mocks base method.
func (m *MockBlock) Message() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(string)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockBlockMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockBlock)(nil).Message))
}

// ModelUUID mocks base method.
func (m *MockBlock) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockBlockMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockBlock)(nil).ModelUUID))
}

// Tag mocks base method.
func (m *MockBlock) Tag() (names.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tag indicates an expected call of Tag.
func (mr *MockBlockMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockBlock)(nil).Tag))
}

// Type mocks base method.
func (m *MockBlock) Type() state.BlockType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(state.BlockType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockBlockMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockBlock)(nil).Type))
}

// updateMessageOp mocks base method.
func (m *MockBlock) updateMessageOp(arg0 string) ([]txn.Op, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "updateMessageOp", arg0)
	ret0, _ := ret[0].([]txn.Op)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// updateMessageOp indicates an expected call of updateMessageOp.
func (mr *MockBlockMockRecorder) updateMessageOp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateMessageOp", reflect.TypeOf((*MockBlock)(nil).updateMessageOp), arg0)
}
