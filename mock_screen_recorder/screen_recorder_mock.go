// Code generated by MockGen. DO NOT EDIT.
// Source: screen_recorder/screen_recorder.go

// Package mock_screen_recorder is a generated GoMock package.
package mock_screen_recorder

import (
	gomock "github.com/golang/mock/gomock"
	output "github.com/ybonjour/atr/output"
	test "github.com/ybonjour/atr/test"
	reflect "reflect"
)

// MockScreenRecorder is a mock of ScreenRecorder interface
type MockScreenRecorder struct {
	ctrl     *gomock.Controller
	recorder *MockScreenRecorderMockRecorder
}

// MockScreenRecorderMockRecorder is the mock recorder for MockScreenRecorder
type MockScreenRecorderMockRecorder struct {
	mock *MockScreenRecorder
}

// NewMockScreenRecorder creates a new mock instance
func NewMockScreenRecorder(ctrl *gomock.Controller) *MockScreenRecorder {
	mock := &MockScreenRecorder{ctrl: ctrl}
	mock.recorder = &MockScreenRecorderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScreenRecorder) EXPECT() *MockScreenRecorderMockRecorder {
	return m.recorder
}

// StartRecording mocks base method
func (m *MockScreenRecorder) StartRecording(test test.Test) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartRecording", test)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartRecording indicates an expected call of StartRecording
func (mr *MockScreenRecorderMockRecorder) StartRecording(test interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRecording", reflect.TypeOf((*MockScreenRecorder)(nil).StartRecording), test)
}

// StopRecording mocks base method
func (m *MockScreenRecorder) StopRecording(test test.Test) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopRecording", test)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopRecording indicates an expected call of StopRecording
func (mr *MockScreenRecorderMockRecorder) StopRecording(test interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopRecording", reflect.TypeOf((*MockScreenRecorder)(nil).StopRecording), test)
}

// SaveResult mocks base method
func (m *MockScreenRecorder) SaveResult(test test.Test, writer output.Writer) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResult", test, writer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveResult indicates an expected call of SaveResult
func (mr *MockScreenRecorderMockRecorder) SaveResult(test, writer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResult", reflect.TypeOf((*MockScreenRecorder)(nil).SaveResult), test, writer)
}

// RemoveRecording mocks base method
func (m *MockScreenRecorder) RemoveRecording(test test.Test) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRecording", test)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRecording indicates an expected call of RemoveRecording
func (mr *MockScreenRecorderMockRecorder) RemoveRecording(test interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRecording", reflect.TypeOf((*MockScreenRecorder)(nil).RemoveRecording), test)
}
