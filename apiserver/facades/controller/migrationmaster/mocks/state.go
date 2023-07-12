// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: ModelMigration,NotifyWatcher)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	migration "github.com/juju/juju/core/migration"
	permission "github.com/juju/juju/core/permission"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockModelMigration is a mock of ModelMigration interface.
type MockModelMigration struct {
	ctrl     *gomock.Controller
	recorder *MockModelMigrationMockRecorder
}

// MockModelMigrationMockRecorder is the mock recorder for MockModelMigration.
type MockModelMigrationMockRecorder struct {
	mock *MockModelMigration
}

// NewMockModelMigration creates a new mock instance.
func NewMockModelMigration(ctrl *gomock.Controller) *MockModelMigration {
	mock := &MockModelMigration{ctrl: ctrl}
	mock.recorder = &MockModelMigrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelMigration) EXPECT() *MockModelMigrationMockRecorder {
	return m.recorder
}

// Attempt mocks base method.
func (m *MockModelMigration) Attempt() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attempt")
	ret0, _ := ret[0].(int)
	return ret0
}

// Attempt indicates an expected call of Attempt.
func (mr *MockModelMigrationMockRecorder) Attempt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attempt", reflect.TypeOf((*MockModelMigration)(nil).Attempt))
}

// EndTime mocks base method.
func (m *MockModelMigration) EndTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EndTime indicates an expected call of EndTime.
func (mr *MockModelMigrationMockRecorder) EndTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndTime", reflect.TypeOf((*MockModelMigration)(nil).EndTime))
}

// Id mocks base method.
func (m *MockModelMigration) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockModelMigrationMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockModelMigration)(nil).Id))
}

// InitiatedBy mocks base method.
func (m *MockModelMigration) InitiatedBy() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiatedBy")
	ret0, _ := ret[0].(string)
	return ret0
}

// InitiatedBy indicates an expected call of InitiatedBy.
func (mr *MockModelMigrationMockRecorder) InitiatedBy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiatedBy", reflect.TypeOf((*MockModelMigration)(nil).InitiatedBy))
}

// MinionReports mocks base method.
func (m *MockModelMigration) MinionReports() (*state.MinionReports, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinionReports")
	ret0, _ := ret[0].(*state.MinionReports)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinionReports indicates an expected call of MinionReports.
func (mr *MockModelMigrationMockRecorder) MinionReports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinionReports", reflect.TypeOf((*MockModelMigration)(nil).MinionReports))
}

// ModelUUID mocks base method.
func (m *MockModelMigration) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockModelMigrationMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockModelMigration)(nil).ModelUUID))
}

// ModelUserAccess mocks base method.
func (m *MockModelMigration) ModelUserAccess(arg0 names.Tag) permission.Access {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUserAccess", arg0)
	ret0, _ := ret[0].(permission.Access)
	return ret0
}

// ModelUserAccess indicates an expected call of ModelUserAccess.
func (mr *MockModelMigrationMockRecorder) ModelUserAccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUserAccess", reflect.TypeOf((*MockModelMigration)(nil).ModelUserAccess), arg0)
}

// Phase mocks base method.
func (m *MockModelMigration) Phase() (migration.Phase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phase")
	ret0, _ := ret[0].(migration.Phase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Phase indicates an expected call of Phase.
func (mr *MockModelMigrationMockRecorder) Phase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phase", reflect.TypeOf((*MockModelMigration)(nil).Phase))
}

// PhaseChangedTime mocks base method.
func (m *MockModelMigration) PhaseChangedTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhaseChangedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// PhaseChangedTime indicates an expected call of PhaseChangedTime.
func (mr *MockModelMigrationMockRecorder) PhaseChangedTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhaseChangedTime", reflect.TypeOf((*MockModelMigration)(nil).PhaseChangedTime))
}

// Refresh mocks base method.
func (m *MockModelMigration) Refresh() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh")
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh.
func (mr *MockModelMigrationMockRecorder) Refresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockModelMigration)(nil).Refresh))
}

// SetPhase mocks base method.
func (m *MockModelMigration) SetPhase(arg0 migration.Phase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPhase", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPhase indicates an expected call of SetPhase.
func (mr *MockModelMigrationMockRecorder) SetPhase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPhase", reflect.TypeOf((*MockModelMigration)(nil).SetPhase), arg0)
}

// SetStatusMessage mocks base method.
func (m *MockModelMigration) SetStatusMessage(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatusMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatusMessage indicates an expected call of SetStatusMessage.
func (mr *MockModelMigrationMockRecorder) SetStatusMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusMessage", reflect.TypeOf((*MockModelMigration)(nil).SetStatusMessage), arg0)
}

// StartTime mocks base method.
func (m *MockModelMigration) StartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// StartTime indicates an expected call of StartTime.
func (mr *MockModelMigrationMockRecorder) StartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTime", reflect.TypeOf((*MockModelMigration)(nil).StartTime))
}

// StatusMessage mocks base method.
func (m *MockModelMigration) StatusMessage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusMessage")
	ret0, _ := ret[0].(string)
	return ret0
}

// StatusMessage indicates an expected call of StatusMessage.
func (mr *MockModelMigrationMockRecorder) StatusMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusMessage", reflect.TypeOf((*MockModelMigration)(nil).StatusMessage))
}

// SubmitMinionReport mocks base method.
func (m *MockModelMigration) SubmitMinionReport(arg0 names.Tag, arg1 migration.Phase, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitMinionReport", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitMinionReport indicates an expected call of SubmitMinionReport.
func (mr *MockModelMigrationMockRecorder) SubmitMinionReport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitMinionReport", reflect.TypeOf((*MockModelMigration)(nil).SubmitMinionReport), arg0, arg1, arg2)
}

// SuccessTime mocks base method.
func (m *MockModelMigration) SuccessTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SuccessTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// SuccessTime indicates an expected call of SuccessTime.
func (mr *MockModelMigrationMockRecorder) SuccessTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SuccessTime", reflect.TypeOf((*MockModelMigration)(nil).SuccessTime))
}

// TargetInfo mocks base method.
func (m *MockModelMigration) TargetInfo() (*migration.TargetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetInfo")
	ret0, _ := ret[0].(*migration.TargetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TargetInfo indicates an expected call of TargetInfo.
func (mr *MockModelMigrationMockRecorder) TargetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetInfo", reflect.TypeOf((*MockModelMigration)(nil).TargetInfo))
}

// WatchMinionReports mocks base method.
func (m *MockModelMigration) WatchMinionReports() (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchMinionReports")
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchMinionReports indicates an expected call of WatchMinionReports.
func (mr *MockModelMigrationMockRecorder) WatchMinionReports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchMinionReports", reflect.TypeOf((*MockModelMigration)(nil).WatchMinionReports))
}

// MockNotifyWatcher is a mock of NotifyWatcher interface.
type MockNotifyWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockNotifyWatcherMockRecorder
}

// MockNotifyWatcherMockRecorder is the mock recorder for MockNotifyWatcher.
type MockNotifyWatcherMockRecorder struct {
	mock *MockNotifyWatcher
}

// NewMockNotifyWatcher creates a new mock instance.
func NewMockNotifyWatcher(ctrl *gomock.Controller) *MockNotifyWatcher {
	mock := &MockNotifyWatcher{ctrl: ctrl}
	mock.recorder = &MockNotifyWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifyWatcher) EXPECT() *MockNotifyWatcherMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockNotifyWatcher) Changes() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockNotifyWatcherMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockNotifyWatcher)(nil).Changes))
}

// Err mocks base method.
func (m *MockNotifyWatcher) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockNotifyWatcherMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockNotifyWatcher)(nil).Err))
}

// Kill mocks base method.
func (m *MockNotifyWatcher) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockNotifyWatcherMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockNotifyWatcher)(nil).Kill))
}

// Stop mocks base method.
func (m *MockNotifyWatcher) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockNotifyWatcherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockNotifyWatcher)(nil).Stop))
}

// Wait mocks base method.
func (m *MockNotifyWatcher) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockNotifyWatcherMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockNotifyWatcher)(nil).Wait))
}
