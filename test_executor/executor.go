package test_executor

import (
	"errors"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/ybonjour/atr/adb"
	"github.com/ybonjour/atr/apks"
	"github.com/ybonjour/atr/devices"
	"github.com/ybonjour/atr/files"
	"github.com/ybonjour/atr/output"
	"github.com/ybonjour/atr/result"
	"github.com/ybonjour/atr/test"
	"github.com/ybonjour/atr/test_listener"
	"sync"
	"time"
)

type Config struct {
	Apk               apks.Apk
	TestApk           apks.Apk
	TestRunner        string
	Tests             []test.Test
	OutputFolder      string
	DisableAnimations bool
}

type Executor interface {
	Execute(config Config, devices []devices.Device) error
}

type executorImpl struct {
	installer     Installer
	resultParser  result.Parser
	adb           adb.Adb
	testListeners []test_listener.TestListener
	writer        output.Writer
	files         files.Files
}

func NewExecutor(writer output.Writer, testListeners []test_listener.TestListener) Executor {
	return executorImpl{
		installer:     NewInstaller(),
		resultParser:  result.NewParser(),
		adb:           adb.New(),
		testListeners: testListeners,
		writer:        writer,
		files:         files.New(),
	}
}

func (executor executorImpl) Execute(config Config, targetDevices []devices.Device) error {
	errorChannel := make(chan error, len(targetDevices))

	var wg sync.WaitGroup
	wg.Add(len(targetDevices))
	for _, targetDevice := range targetDevices {
		go func(d devices.Device) {
			err := executor.executeOnDevice(config, d)
			if err != nil {
				errorChannel <- err
			}

			wg.Done()
		}(targetDevice)
	}
	go func() {
		wg.Wait()
		close(errorChannel)
	}()

	var collectedErrors error
	for err := range errorChannel {
		collectedErrors = multierror.Append(collectedErrors, err)
	}

	return collectedErrors
}

func (executor executorImpl) executeOnDevice(config Config, device devices.Device) error {
	installError := executor.reinstallApks(config, device)
	if installError != nil {
		return installError
	}
	directory, directoryError := executor.writer.GetDeviceDirectory(device)
	if directoryError != nil {
		return directoryError
	}
	removeError := executor.files.RemoveDirectory(directory)
	if removeError != nil {
		return removeError
	}
	if config.DisableAnimations {
		disableAnimationsError := executor.adb.DisableAnimations(device.Serial)
		if disableAnimationsError != nil {
			return disableAnimationsError
		}
	}

	return executor.executeTests(config, device)
}

func (executor executorImpl) reinstallApks(config Config, device devices.Device) error {
	apkInstallError := executor.installer.Reinstall(config.Apk, device)
	if apkInstallError != nil {
		return apkInstallError
	}
	testApkInstallError := executor.installer.Reinstall(config.TestApk, device)
	if testApkInstallError != nil {
		return testApkInstallError
	}

	return nil
}

func (executor executorImpl) executeTests(testConfig Config, device devices.Device) error {
	executor.beforeTestSuite(device)
	var testSuiteResult error
	var results []result.Result
	for _, t := range testConfig.Tests {
		executor.beforeTest(t)
		testOutput, errTest, duration := executor.executeSingleTest(t, device, testConfig.TestApk.PackageName, testConfig.TestRunner)
		r := executor.resultParser.ParseFromOutput(t, errTest, testOutput, duration)
		if r.IsFailure() {
			testSuiteResult = multierror.Append(testSuiteResult, errors.New(fmt.Sprintf("Test '%v' failed on device '%v'", r.Test.FullName(), device)))
		}
		extendedResult := executor.afterTest(r)
		results = append(results, extendedResult)

	}
	executor.afterTestSuite()

	return testSuiteResult
}

func (executor executorImpl) forAllTestListeners(f func(listener test_listener.TestListener)) {
	for _, listener := range executor.testListeners {
		f(listener)
	}
}

func (executor executorImpl) beforeTestSuite(device devices.Device) {
	executor.forAllTestListeners(func(listener test_listener.TestListener) {
		listener.BeforeTestSuite(device)
	})
}

func (executor executorImpl) afterTestSuite() {
	executor.forAllTestListeners(func(listener test_listener.TestListener) {
		listener.AfterTestSuite()
	})
}

func (executor executorImpl) beforeTest(t test.Test) {
	executor.forAllTestListeners(func(listener test_listener.TestListener) {
		listener.BeforeTest(t)
	})
}

func (executor executorImpl) afterTest(r result.Result) result.Result {
	executor.forAllTestListeners(func(listener test_listener.TestListener) {
		extras := listener.AfterTest(r)
		r.Extras = append(r.Extras, extras...)
	})

	return r
}

func (executor executorImpl) executeSingleTest(t test.Test, device devices.Device, testPackage string, testRunner string) (string, error, time.Duration) {
	start := time.Now()
	testOutput, err := executor.adb.ExecuteTest(testPackage, testRunner, t.FullName(), device.Serial)
	duration := time.Since(start)
	return testOutput, err, duration
}
