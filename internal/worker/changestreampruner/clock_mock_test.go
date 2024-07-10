// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/clock (interfaces: Clock,Timer)
//
// Generated by this command:
//
//	mockgen -typed -package changestreampruner -destination clock_mock_test.go github.com/juju/clock Clock,Timer
//

// Package changestreampruner is a generated GoMock package.
package changestreampruner

import (
	reflect "reflect"
	time "time"

	clock "github.com/juju/clock"
	gomock "go.uber.org/mock/gomock"
)

// MockClock is a mock of Clock interface.
type MockClock struct {
	ctrl     *gomock.Controller
	recorder *MockClockMockRecorder
}

// MockClockMockRecorder is the mock recorder for MockClock.
type MockClockMockRecorder struct {
	mock *MockClock
}

// NewMockClock creates a new mock instance.
func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &MockClockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClock) EXPECT() *MockClockMockRecorder {
	return m.recorder
}

// After mocks base method.
func (m *MockClock) After(arg0 time.Duration) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "After", arg0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// After indicates an expected call of After.
func (mr *MockClockMockRecorder) After(arg0 any) *MockClockAfterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "After", reflect.TypeOf((*MockClock)(nil).After), arg0)
	return &MockClockAfterCall{Call: call}
}

// MockClockAfterCall wrap *gomock.Call
type MockClockAfterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAfterCall) Return(arg0 <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAfterCall) Do(f func(time.Duration) <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAfterCall) DoAndReturn(f func(time.Duration) <-chan time.Time) *MockClockAfterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AfterFunc mocks base method.
func (m *MockClock) AfterFunc(arg0 time.Duration, arg1 func()) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFunc", arg0, arg1)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc.
func (mr *MockClockMockRecorder) AfterFunc(arg0, arg1 any) *MockClockAfterFuncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockClock)(nil).AfterFunc), arg0, arg1)
	return &MockClockAfterFuncCall{Call: call}
}

// MockClockAfterFuncCall wrap *gomock.Call
type MockClockAfterFuncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAfterFuncCall) Return(arg0 clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAfterFuncCall) Do(f func(time.Duration, func()) clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAfterFuncCall) DoAndReturn(f func(time.Duration, func()) clock.Timer) *MockClockAfterFuncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// At mocks base method.
func (m *MockClock) At(arg0 time.Time) <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "At", arg0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// At indicates an expected call of At.
func (mr *MockClockMockRecorder) At(arg0 any) *MockClockAtCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "At", reflect.TypeOf((*MockClock)(nil).At), arg0)
	return &MockClockAtCall{Call: call}
}

// MockClockAtCall wrap *gomock.Call
type MockClockAtCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAtCall) Return(arg0 <-chan time.Time) *MockClockAtCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAtCall) Do(f func(time.Time) <-chan time.Time) *MockClockAtCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAtCall) DoAndReturn(f func(time.Time) <-chan time.Time) *MockClockAtCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AtFunc mocks base method.
func (m *MockClock) AtFunc(arg0 time.Time, arg1 func()) clock.Alarm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AtFunc", arg0, arg1)
	ret0, _ := ret[0].(clock.Alarm)
	return ret0
}

// AtFunc indicates an expected call of AtFunc.
func (mr *MockClockMockRecorder) AtFunc(arg0, arg1 any) *MockClockAtFuncCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AtFunc", reflect.TypeOf((*MockClock)(nil).AtFunc), arg0, arg1)
	return &MockClockAtFuncCall{Call: call}
}

// MockClockAtFuncCall wrap *gomock.Call
type MockClockAtFuncCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockAtFuncCall) Return(arg0 clock.Alarm) *MockClockAtFuncCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockAtFuncCall) Do(f func(time.Time, func()) clock.Alarm) *MockClockAtFuncCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockAtFuncCall) DoAndReturn(f func(time.Time, func()) clock.Alarm) *MockClockAtFuncCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewAlarm mocks base method.
func (m *MockClock) NewAlarm(arg0 time.Time) clock.Alarm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAlarm", arg0)
	ret0, _ := ret[0].(clock.Alarm)
	return ret0
}

// NewAlarm indicates an expected call of NewAlarm.
func (mr *MockClockMockRecorder) NewAlarm(arg0 any) *MockClockNewAlarmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAlarm", reflect.TypeOf((*MockClock)(nil).NewAlarm), arg0)
	return &MockClockNewAlarmCall{Call: call}
}

// MockClockNewAlarmCall wrap *gomock.Call
type MockClockNewAlarmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockNewAlarmCall) Return(arg0 clock.Alarm) *MockClockNewAlarmCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockNewAlarmCall) Do(f func(time.Time) clock.Alarm) *MockClockNewAlarmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockNewAlarmCall) DoAndReturn(f func(time.Time) clock.Alarm) *MockClockNewAlarmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NewTimer mocks base method.
func (m *MockClock) NewTimer(arg0 time.Duration) clock.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", arg0)
	ret0, _ := ret[0].(clock.Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer.
func (mr *MockClockMockRecorder) NewTimer(arg0 any) *MockClockNewTimerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockClock)(nil).NewTimer), arg0)
	return &MockClockNewTimerCall{Call: call}
}

// MockClockNewTimerCall wrap *gomock.Call
type MockClockNewTimerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockNewTimerCall) Return(arg0 clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockNewTimerCall) Do(f func(time.Duration) clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockNewTimerCall) DoAndReturn(f func(time.Duration) clock.Timer) *MockClockNewTimerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Now mocks base method.
func (m *MockClock) Now() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Now indicates an expected call of Now.
func (mr *MockClockMockRecorder) Now() *MockClockNowCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Now", reflect.TypeOf((*MockClock)(nil).Now))
	return &MockClockNowCall{Call: call}
}

// MockClockNowCall wrap *gomock.Call
type MockClockNowCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockClockNowCall) Return(arg0 time.Time) *MockClockNowCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockClockNowCall) Do(f func() time.Time) *MockClockNowCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockClockNowCall) DoAndReturn(f func() time.Time) *MockClockNowCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockTimer is a mock of Timer interface.
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer.
type MockTimerMockRecorder struct {
	mock *MockTimer
}

// NewMockTimer creates a new mock instance.
func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &MockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimer) EXPECT() *MockTimerMockRecorder {
	return m.recorder
}

// Chan mocks base method.
func (m *MockTimer) Chan() <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chan")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// Chan indicates an expected call of Chan.
func (mr *MockTimerMockRecorder) Chan() *MockTimerChanCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chan", reflect.TypeOf((*MockTimer)(nil).Chan))
	return &MockTimerChanCall{Call: call}
}

// MockTimerChanCall wrap *gomock.Call
type MockTimerChanCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTimerChanCall) Return(arg0 <-chan time.Time) *MockTimerChanCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTimerChanCall) Do(f func() <-chan time.Time) *MockTimerChanCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTimerChanCall) DoAndReturn(f func() <-chan time.Time) *MockTimerChanCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Reset mocks base method.
func (m *MockTimer) Reset(arg0 time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockTimerMockRecorder) Reset(arg0 any) *MockTimerResetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTimer)(nil).Reset), arg0)
	return &MockTimerResetCall{Call: call}
}

// MockTimerResetCall wrap *gomock.Call
type MockTimerResetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTimerResetCall) Return(arg0 bool) *MockTimerResetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTimerResetCall) Do(f func(time.Duration) bool) *MockTimerResetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTimerResetCall) DoAndReturn(f func(time.Duration) bool) *MockTimerResetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *MockTimer) Stop() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockTimerMockRecorder) Stop() *MockTimerStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTimer)(nil).Stop))
	return &MockTimerStopCall{Call: call}
}

// MockTimerStopCall wrap *gomock.Call
type MockTimerStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTimerStopCall) Return(arg0 bool) *MockTimerStopCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTimerStopCall) Do(f func() bool) *MockTimerStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTimerStopCall) DoAndReturn(f func() bool) *MockTimerStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
