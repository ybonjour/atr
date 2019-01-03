package test_executor

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/ybonjour/atr/apks"
	"github.com/ybonjour/atr/devices"
	"github.com/ybonjour/atr/logcat"
	"github.com/ybonjour/atr/mock_adb"
	"github.com/ybonjour/atr/mock_logcat"
	"github.com/ybonjour/atr/mock_result"
	"github.com/ybonjour/atr/mock_test_executor"
	"github.com/ybonjour/atr/result"
	"github.com/ybonjour/atr/test"
	"testing"
)

func TestExecute(t *testing.T) {
	targetTest := test.Test{Class: "TestClass", Method: "testMethod"}
	config := Config{
		TestApk:    apks.Apk{PackageName: "testPackageName"},
		Tests:      []test.Test{targetTest},
		TestRunner: "testRunner",
	}
	testOutput := "testOutput"
	testResult := result.Result{}
	device := devices.Device{Serial: "abcd"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInstaller := mock_test_executor.NewMockInstaller(ctrl)
	mockInstaller.EXPECT().Reinstall(config.Apk, device).Return(nil)
	mockInstaller.EXPECT().Reinstall(config.TestApk, device).Return(nil)
	mockAdb := mock_adb.NewMockAdb(ctrl)
	mockAdb.
		EXPECT().
		ExecuteTest(config.TestApk.PackageName, config.TestRunner, targetTest.FullName(), device.Serial).
		Return(testOutput, nil)
	mockResultParser := mock_result.NewMockParser(ctrl)
	mockResultParser.EXPECT().ParseFromOutput(gomock.Eq(targetTest), gomock.Eq(nil), gomock.Eq(testOutput), gomock.Any()).Return(testResult)
	executor := executorImpl{
		installer:     mockInstaller,
		adb:           mockAdb,
		resultParser:  mockResultParser,
		logcatFactory: mockLogcatFactory(map[devices.Device][]test.Test{device: {targetTest}}, ctrl),
	}

	results, err := executor.Execute(config, []devices.Device{device})

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}
	expectedResults := []result.Result{testResult}
	if !AreEqualResults(results[device], expectedResults) {
		t.Error(fmt.Sprintf("Expected results '%v' but got '%v'", expectedResults, results[device]))
	}
}

func TestExecuteMultipleTests(t *testing.T) {
	test1 := test.Test{Class: "TestClass", Method: "testMethod"}
	test2 := test.Test{Class: "TestClass", Method: "testMethod1"}
	testResult1 := result.Result{}
	testResult2 := result.Result{}
	device := devices.Device{Serial: "abcd"}
	config := Config{
		Tests: []test.Test{test1, test2},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInstaller := mock_test_executor.NewMockInstaller(ctrl)
	mockAdb := mock_adb.NewMockAdb(ctrl)
	mockResultParser := mock_result.NewMockParser(ctrl)
	givenAllApksInstalledSuccessfully(mockInstaller, 1)
	givenTestOnDeviceReturns(test1, device, testResult1, mockAdb, mockResultParser)
	givenTestOnDeviceReturns(test2, device, testResult2, mockAdb, mockResultParser)
	executor := executorImpl{
		installer:     mockInstaller,
		adb:           mockAdb,
		resultParser:  mockResultParser,
		logcatFactory: mockLogcatFactory(map[devices.Device][]test.Test{device: {test1, test2}}, ctrl),
	}

	results, err := executor.Execute(config, []devices.Device{device})

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}
	expectedResults := []result.Result{testResult1, testResult2}
	if !AreEqualResults(results[device], expectedResults) {
		t.Error(fmt.Sprintf("Expected results '%v' but got '%v'", expectedResults, results[device]))
	}
}

func TestExecuteMultipleDevices(t *testing.T) {
	targetTest := test.Test{Class: "TestClass", Method: "testMethod"}
	testResult1 := result.Result{}
	testResult2 := result.Result{}
	device1 := devices.Device{Serial: "abcd"}
	device2 := devices.Device{Serial: "efgh"}
	config := Config{
		Tests: []test.Test{targetTest},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockInstaller := mock_test_executor.NewMockInstaller(ctrl)
	mockAdb := mock_adb.NewMockAdb(ctrl)
	mockResultParser := mock_result.NewMockParser(ctrl)
	givenAllApksInstalledSuccessfully(mockInstaller, 2)
	givenTestOnDeviceReturns(targetTest, device1, testResult1, mockAdb, mockResultParser)
	givenTestOnDeviceReturns(targetTest, device2, testResult2, mockAdb, mockResultParser)
	executor := executorImpl{
		installer:     mockInstaller,
		adb:           mockAdb,
		resultParser:  mockResultParser,
		logcatFactory: mockLogcatFactory(map[devices.Device][]test.Test{device1: {targetTest}, device2: {targetTest}}, ctrl),
	}

	results, err := executor.Execute(config, []devices.Device{device1, device2})

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}
	expectedResultsDevice1 := []result.Result{testResult1}
	if !AreEqualResults(results[device1], expectedResultsDevice1) {
		t.Error(fmt.Sprintf("Expected results '%v' but got '%v'", expectedResultsDevice1, results[device1]))
	}
	expectedResultsDevice2 := []result.Result{testResult2}
	if !AreEqualResults(results[device1], expectedResultsDevice2) {
		t.Error(fmt.Sprintf("Expected results '%v' but got '%v'", expectedResultsDevice2, results[device2]))
	}
}

func givenAllApksInstalledSuccessfully(mockInstaller *mock_test_executor.MockInstaller, numDevices int) {
	mockInstaller.EXPECT().Reinstall(gomock.Any(), gomock.Any()).Return(nil).Times(numDevices * 2)
}

func givenTestOnDeviceReturns(t test.Test, d devices.Device, r result.Result, mockAdb *mock_adb.MockAdb, mockResultParser *mock_result.MockParser) {
	testOutput := t.FullName()
	mockAdb.
		EXPECT().
		ExecuteTest(gomock.Any(), gomock.Any(), gomock.Eq(t.FullName()), gomock.Eq(d.Serial)).
		Return(testOutput, nil)

	mockResultParser.
		EXPECT().
		ParseFromOutput(gomock.Eq(t), gomock.Eq(nil), gomock.Eq(testOutput), gomock.Any()).
		Return(r)
}

func mockLogcatFactory(testsPerDevice map[devices.Device][]test.Test, ctrl *gomock.Controller) logcat.Factory {
	mockLogcatFactory := mock_logcat.NewMockFactory(ctrl)
	for device, tests := range testsPerDevice {
		mockLogcat := mockLogcat(tests, ctrl)
		mockLogcatFactory.EXPECT().ForDevice(device).Return(mockLogcat)
	}

	return mockLogcatFactory
}

func mockLogcat(tests []test.Test, ctrl *gomock.Controller) logcat.Logcat {
	mockLogcat := mock_logcat.NewMockLogcat(ctrl)
	for _, t := range tests {
		mockLogcat.EXPECT().StartRecording(t).Return(nil)
		mockLogcat.EXPECT().StopRecording(t).Return(nil)
	}

	return mockLogcat
}

func AreEqualResults(slice1, slice2 []result.Result) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
