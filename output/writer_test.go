package output

import (
	"fmt"
	"github.com/android-test-runner/atr/devices"
	"github.com/android-test-runner/atr/files"
	"github.com/android-test-runner/atr/mock_files"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestWriteFile(t *testing.T) {
	rootDir := "rootDir"
	file := files.File{Name: "filename", Content: "content"}
	device := devices.Device{Serial: "deviceSerial"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filesMock := mock_files.NewMockFiles(ctrl)
	filesMock.EXPECT().WriteFile("rootDir/deviceSerial/filename", file.Content).Return(nil)

	writer := writerImpl{
		rootDir: rootDir,
		files:   filesMock,
	}

	filepath, err := writer.WriteFile(file, device)

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}

	if filepath != "deviceSerial/filename" {
		t.Error(fmt.Sprintf("Expected filepath 'device/filename' but got '%v'", filepath))
	}
}

func TestWriteFileToRoot(t *testing.T) {
	rootDir := "rootDir"

	file := files.File{Name: "filename", Content: "content"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	filesMock := mock_files.NewMockFiles(ctrl)
	filesMock.EXPECT().WriteFile("rootDir/filename", file.Content).Return(nil)

	writer := writerImpl{
		rootDir: rootDir,
		files:   filesMock,
	}

	filepath, err := writer.WriteFileToRoot(file)

	if err != nil {
		t.Error(fmt.Sprintf("Expected no error but got '%v'", err))
	}

	if filepath != "filename" {
		t.Error(fmt.Sprintf("Expected filepath 'filename' but got '%v'", filepath))
	}
}
