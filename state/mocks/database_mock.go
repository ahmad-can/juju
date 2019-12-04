// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: Database)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongo "github.com/juju/juju/mongo"
	state "github.com/juju/juju/state"
	txn "github.com/juju/txn"
	mgo_v2 "gopkg.in/mgo.v2"
	txn0 "gopkg.in/mgo.v2/txn"
	reflect "reflect"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// Copy mocks base method
func (m *MockDatabase) Copy() (state.Database, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(state.Database)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// Copy indicates an expected call of Copy
func (mr *MockDatabaseMockRecorder) Copy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockDatabase)(nil).Copy))
}

// CopyForModel mocks base method
func (m *MockDatabase) CopyForModel(arg0 string) (state.Database, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyForModel", arg0)
	ret0, _ := ret[0].(state.Database)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// CopyForModel indicates an expected call of CopyForModel
func (mr *MockDatabaseMockRecorder) CopyForModel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyForModel", reflect.TypeOf((*MockDatabase)(nil).CopyForModel), arg0)
}

// GetCollection mocks base method
func (m *MockDatabase) GetCollection(arg0 string) (mongo.Collection, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", arg0)
	ret0, _ := ret[0].(mongo.Collection)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection
func (mr *MockDatabaseMockRecorder) GetCollection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockDatabase)(nil).GetCollection), arg0)
}

// GetCollectionFor mocks base method
func (m *MockDatabase) GetCollectionFor(arg0, arg1 string) (mongo.Collection, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionFor", arg0, arg1)
	ret0, _ := ret[0].(mongo.Collection)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// GetCollectionFor indicates an expected call of GetCollectionFor
func (mr *MockDatabaseMockRecorder) GetCollectionFor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionFor", reflect.TypeOf((*MockDatabase)(nil).GetCollectionFor), arg0, arg1)
}

// GetRawCollection mocks base method
func (m *MockDatabase) GetRawCollection(arg0 string) (*mgo_v2.Collection, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawCollection", arg0)
	ret0, _ := ret[0].(*mgo_v2.Collection)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// GetRawCollection indicates an expected call of GetRawCollection
func (mr *MockDatabaseMockRecorder) GetRawCollection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawCollection", reflect.TypeOf((*MockDatabase)(nil).GetRawCollection), arg0)
}

// Run mocks base method
func (m *MockDatabase) Run(arg0 txn.TransactionSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockDatabaseMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockDatabase)(nil).Run), arg0)
}

// RunRawTransaction mocks base method
func (m *MockDatabase) RunRawTransaction(arg0 []txn0.Op) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunRawTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunRawTransaction indicates an expected call of RunRawTransaction
func (mr *MockDatabaseMockRecorder) RunRawTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunRawTransaction", reflect.TypeOf((*MockDatabase)(nil).RunRawTransaction), arg0)
}

// RunTransaction mocks base method
func (m *MockDatabase) RunTransaction(arg0 []txn0.Op) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction
func (mr *MockDatabaseMockRecorder) RunTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockDatabase)(nil).RunTransaction), arg0)
}

// RunTransactionFor mocks base method
func (m *MockDatabase) RunTransactionFor(arg0 string, arg1 []txn0.Op) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransactionFor", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransactionFor indicates an expected call of RunTransactionFor
func (mr *MockDatabaseMockRecorder) RunTransactionFor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransactionFor", reflect.TypeOf((*MockDatabase)(nil).RunTransactionFor), arg0, arg1)
}

// Schema mocks base method
func (m *MockDatabase) Schema() state.CollectionSchema {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Schema")
	ret0, _ := ret[0].(state.CollectionSchema)
	return ret0
}

// Schema indicates an expected call of Schema
func (mr *MockDatabaseMockRecorder) Schema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Schema", reflect.TypeOf((*MockDatabase)(nil).Schema))
}

// TransactionRunner mocks base method
func (m *MockDatabase) TransactionRunner() (txn.Runner, state.SessionCloser) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionRunner")
	ret0, _ := ret[0].(txn.Runner)
	ret1, _ := ret[1].(state.SessionCloser)
	return ret0, ret1
}

// TransactionRunner indicates an expected call of TransactionRunner
func (mr *MockDatabaseMockRecorder) TransactionRunner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionRunner", reflect.TypeOf((*MockDatabase)(nil).TransactionRunner))
}
