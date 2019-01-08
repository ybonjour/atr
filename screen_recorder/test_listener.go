package screen_recorder

import (
	"fmt"
	"github.com/ybonjour/atr/devices"
	"github.com/ybonjour/atr/output"
	"github.com/ybonjour/atr/result"
	"github.com/ybonjour/atr/test"
	"github.com/ybonjour/atr/test_listener"
)

type testListener struct {
	writer         output.Writer
	screenRecorder ScreenRecorder
}

func NewTestListener(writer output.Writer) test_listener.TestListener {
	return &testListener{
		writer: writer,
	}
}

func (listener *testListener) BeforeTestSuite(device devices.Device) {
	listener.screenRecorder = New(device)
}

func (listener *testListener) AfterTestSuite() {}

func (listener *testListener) BeforeTest(test test.Test) {
	errStartScreenRecording := listener.screenRecorder.StartRecording(test)
	if errStartScreenRecording != nil {
		fmt.Printf("Could not start screen recording: '%v'\n", errStartScreenRecording)
	}
}

func (listener *testListener) AfterTest(r result.Result) []result.Extra {
	errStopScreenRecording := listener.screenRecorder.StopRecording(r.Test)
	if errStopScreenRecording != nil {
		fmt.Printf("Could not save screen recording: '%v'\n", errStopScreenRecording)
	}

	if r.IsFailure() {
		errSave := listener.screenRecorder.SaveResult(r.Test, listener.writer)
		if errSave != nil {
			fmt.Printf("Could not save screen recording: '%v'\n", errSave)
		}
	}

	errRemove := listener.screenRecorder.RemoveRecording(r.Test)

	if errRemove != nil {
		fmt.Printf("Could not remove screen recording: '%v'\n", errRemove)
	}

	return []result.Extra{}
}
