package test_executor

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/ybonjour/atr/apks"
	"github.com/ybonjour/atr/devices"
	"github.com/ybonjour/atr/mock_adb"
	"github.com/ybonjour/atr/mock_result"
	"github.com/ybonjour/atr/mock_test_executor"
	"github.com/ybonjour/atr/result"
	"github.com/ybonjour/atr/test"
	"testing"
)

func TestExecute(t *testing.T) {
	targetTest := test.Test{Class: "TestClass", Method: "testMethod"}
	config := Config{
		Apk:        apks.Apk{},
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
	mockResultParser := mock_result.NewMockResultParser(ctrl)
	mockResultParser.EXPECT().ParseFromOutput(targetTest, nil, testOutput).Return(testResult)
	executor := executorImpl{
		installer:    mockInstaller,
		adb:          mockAdb,
		resultParser: mockResultParser,
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
	givenAllApksInstalledSuccessfully(mockInstaller)
	mockResultParser := mock_result.NewMockResultParser(ctrl)
	givenTestReturns(test1, testResult1, mockAdb, mockResultParser)
	givenTestReturns(test2, testResult2, mockAdb, mockResultParser)
	executor := executorImpl{
		installer:    mockInstaller,
		adb:          mockAdb,
		resultParser: mockResultParser,
	}

	results, err := executor.Execute(config, []devices.Device{device})

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}
	expectedResults := []result.Result{testResult1, testResult2}
	if !AreEqualResults(results[device], expectedResults) {
		t.Error(fmt.Sprintf("Expected results '%v' but got '%v'", expectedResults, results[devices.Device{Serial: "abcd"}]))
	}
}

func givenAllApksInstalledSuccessfully(mockInstaller *mock_test_executor.MockInstaller) {
	mockInstaller.EXPECT().Reinstall(gomock.Any(), gomock.Any()).Return(nil).Times(2)
}

func givenTestReturns(t test.Test, r result.Result, mockAdb *mock_adb.MockAdb, mockResultParser *mock_result.MockResultParser) {
	testOutput := t.FullName()
	mockAdb.
		EXPECT().
		ExecuteTest(gomock.Any(), gomock.Any(), gomock.Eq(t.FullName()), gomock.Any()).
		Return(testOutput, nil)

	mockResultParser.
		EXPECT().
		ParseFromOutput(t, nil, testOutput).
		Return(r)
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
