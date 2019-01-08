// Code generated by MockGen. DO NOT EDIT.
// Source: test_listener/test_listener.go

// Package mock_test_listener is a generated GoMock package.
package mock_test_listener

import (
	gomock "github.com/golang/mock/gomock"
	devices "github.com/ybonjour/atr/devices"
	result "github.com/ybonjour/atr/result"
	test "github.com/ybonjour/atr/test"
	reflect "reflect"
)

// MockTestListener is a mock of TestListener interface
type MockTestListener struct {
	ctrl     *gomock.Controller
	recorder *MockTestListenerMockRecorder
}

// MockTestListenerMockRecorder is the mock recorder for MockTestListener
type MockTestListenerMockRecorder struct {
	mock *MockTestListener
}

// NewMockTestListener creates a new mock instance
func NewMockTestListener(ctrl *gomock.Controller) *MockTestListener {
	mock := &MockTestListener{ctrl: ctrl}
	mock.recorder = &MockTestListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTestListener) EXPECT() *MockTestListenerMockRecorder {
	return m.recorder
}

// BeforeTestSuite mocks base method
func (m *MockTestListener) BeforeTestSuite(device devices.Device) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeTestSuite", device)
}

// BeforeTestSuite indicates an expected call of BeforeTestSuite
func (mr *MockTestListenerMockRecorder) BeforeTestSuite(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeTestSuite", reflect.TypeOf((*MockTestListener)(nil).BeforeTestSuite), device)
}

// AfterTestSuite mocks base method
func (m *MockTestListener) AfterTestSuite() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterTestSuite")
}

// AfterTestSuite indicates an expected call of AfterTestSuite
func (mr *MockTestListenerMockRecorder) AfterTestSuite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterTestSuite", reflect.TypeOf((*MockTestListener)(nil).AfterTestSuite))
}

// BeforeTest mocks base method
func (m *MockTestListener) BeforeTest(test test.Test) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeforeTest", test)
}

// BeforeTest indicates an expected call of BeforeTest
func (mr *MockTestListenerMockRecorder) BeforeTest(test interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeTest", reflect.TypeOf((*MockTestListener)(nil).BeforeTest), test)
}

// AfterTest mocks base method
func (m *MockTestListener) AfterTest(r result.Result) []result.Extra {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterTest", r)
	ret0, _ := ret[0].([]result.Extra)
	return ret0
}

// AfterTest indicates an expected call of AfterTest
func (mr *MockTestListenerMockRecorder) AfterTest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterTest", reflect.TypeOf((*MockTestListener)(nil).AfterTest), r)
}
