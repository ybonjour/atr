// Code generated by MockGen. DO NOT EDIT.
// Source: adb/output_parser.go

// Package mock_adb is a generated GoMock package.
package mock_adb

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockoutputParser is a mock of outputParser interface
type MockoutputParser struct {
	ctrl     *gomock.Controller
	recorder *MockoutputParserMockRecorder
}

// MockoutputParserMockRecorder is the mock recorder for MockoutputParser
type MockoutputParserMockRecorder struct {
	mock *MockoutputParser
}

// NewMockoutputParser creates a new mock instance
func NewMockoutputParser(ctrl *gomock.Controller) *MockoutputParser {
	mock := &MockoutputParser{ctrl: ctrl}
	mock.recorder = &MockoutputParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockoutputParser) EXPECT() *MockoutputParserMockRecorder {
	return m.recorder
}

// ParseConnectedDeviceSerials mocks base method
func (m *MockoutputParser) ParseConnectedDeviceSerials(out string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseConnectedDeviceSerials", out)
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParseConnectedDeviceSerials indicates an expected call of ParseConnectedDeviceSerials
func (mr *MockoutputParserMockRecorder) ParseConnectedDeviceSerials(out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseConnectedDeviceSerials", reflect.TypeOf((*MockoutputParser)(nil).ParseConnectedDeviceSerials), out)
}
