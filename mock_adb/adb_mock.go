// Code generated by MockGen. DO NOT EDIT.
// Source: adb/adb.go

// Package mock_adb is a generated GoMock package.
package mock_adb

import (
	gomock "github.com/golang/mock/gomock"
	command "github.com/ybonjour/atr/command"
	reflect "reflect"
)

// MockAdb is a mock of Adb interface
type MockAdb struct {
	ctrl     *gomock.Controller
	recorder *MockAdbMockRecorder
}

// MockAdbMockRecorder is the mock recorder for MockAdb
type MockAdbMockRecorder struct {
	mock *MockAdb
}

// NewMockAdb creates a new mock instance
func NewMockAdb(ctrl *gomock.Controller) *MockAdb {
	mock := &MockAdb{ctrl: ctrl}
	mock.recorder = &MockAdbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdb) EXPECT() *MockAdbMockRecorder {
	return m.recorder
}

// Version mocks base method
func (m *MockAdb) Version() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockAdbMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockAdb)(nil).Version))
}

// ConnectedDevices mocks base method
func (m *MockAdb) ConnectedDevices() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedDevices")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectedDevices indicates an expected call of ConnectedDevices
func (mr *MockAdbMockRecorder) ConnectedDevices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedDevices", reflect.TypeOf((*MockAdb)(nil).ConnectedDevices))
}

// DisableAnimations mocks base method
func (m *MockAdb) DisableAnimations(deviceSerial string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableAnimations", deviceSerial)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableAnimations indicates an expected call of DisableAnimations
func (mr *MockAdbMockRecorder) DisableAnimations(deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableAnimations", reflect.TypeOf((*MockAdb)(nil).DisableAnimations), deviceSerial)
}

// Install mocks base method
func (m *MockAdb) Install(apkPath, deviceSerial string) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", apkPath, deviceSerial)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockAdbMockRecorder) Install(apkPath, deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAdb)(nil).Install), apkPath, deviceSerial)
}

// Uninstall mocks base method
func (m *MockAdb) Uninstall(packageName, deviceSerial string) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", packageName, deviceSerial)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// Uninstall indicates an expected call of Uninstall
func (mr *MockAdbMockRecorder) Uninstall(packageName, deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockAdb)(nil).Uninstall), packageName, deviceSerial)
}

// ExecuteTest mocks base method
func (m *MockAdb) ExecuteTest(packageName, testRunner, test, deviceSerial string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteTest", packageName, testRunner, test, deviceSerial)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTest indicates an expected call of ExecuteTest
func (mr *MockAdbMockRecorder) ExecuteTest(packageName, testRunner, test, deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTest", reflect.TypeOf((*MockAdb)(nil).ExecuteTest), packageName, testRunner, test, deviceSerial)
}

// ClearLogcat mocks base method
func (m *MockAdb) ClearLogcat(deviceSerial string) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearLogcat", deviceSerial)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// ClearLogcat indicates an expected call of ClearLogcat
func (mr *MockAdbMockRecorder) ClearLogcat(deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearLogcat", reflect.TypeOf((*MockAdb)(nil).ClearLogcat), deviceSerial)
}

// GetLogcat mocks base method
func (m *MockAdb) GetLogcat(deviceSerial string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogcat", deviceSerial)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogcat indicates an expected call of GetLogcat
func (mr *MockAdbMockRecorder) GetLogcat(deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogcat", reflect.TypeOf((*MockAdb)(nil).GetLogcat), deviceSerial)
}

// RecordScreenInBackground mocks base method
func (m *MockAdb) RecordScreenInBackground(deviceSerial, filePath, screenDimensions string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordScreenInBackground", deviceSerial, filePath, screenDimensions)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordScreenInBackground indicates an expected call of RecordScreenInBackground
func (mr *MockAdbMockRecorder) RecordScreenInBackground(deviceSerial, filePath, screenDimensions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordScreenInBackground", reflect.TypeOf((*MockAdb)(nil).RecordScreenInBackground), deviceSerial, filePath, screenDimensions)
}

// RecordScreen mocks base method
func (m *MockAdb) RecordScreen(deviceSerial, filePath, screenDimensions string, timeLimitSeconds int) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordScreen", deviceSerial, filePath, screenDimensions, timeLimitSeconds)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// RecordScreen indicates an expected call of RecordScreen
func (mr *MockAdbMockRecorder) RecordScreen(deviceSerial, filePath, screenDimensions, timeLimitSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordScreen", reflect.TypeOf((*MockAdb)(nil).RecordScreen), deviceSerial, filePath, screenDimensions, timeLimitSeconds)
}

// GetScreenDimensions mocks base method
func (m *MockAdb) GetScreenDimensions(deviceSerial string) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenDimensions", deviceSerial)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetScreenDimensions indicates an expected call of GetScreenDimensions
func (mr *MockAdbMockRecorder) GetScreenDimensions(deviceSerial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenDimensions", reflect.TypeOf((*MockAdb)(nil).GetScreenDimensions), deviceSerial)
}

// PullFile mocks base method
func (m *MockAdb) PullFile(deviceSerial, filePathOnDevice, filePathLocal string) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullFile", deviceSerial, filePathOnDevice, filePathLocal)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// PullFile indicates an expected call of PullFile
func (mr *MockAdbMockRecorder) PullFile(deviceSerial, filePathOnDevice, filePathLocal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullFile", reflect.TypeOf((*MockAdb)(nil).PullFile), deviceSerial, filePathOnDevice, filePathLocal)
}

// RemoveFile mocks base method
func (m *MockAdb) RemoveFile(deviceSerial, filePathOnDevice string) command.ExecutionResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", deviceSerial, filePathOnDevice)
	ret0, _ := ret[0].(command.ExecutionResult)
	return ret0
}

// RemoveFile indicates an expected call of RemoveFile
func (mr *MockAdbMockRecorder) RemoveFile(deviceSerial, filePathOnDevice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MockAdb)(nil).RemoveFile), deviceSerial, filePathOnDevice)
}
