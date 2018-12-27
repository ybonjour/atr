package device

import (
	"github.com/ybonjour/atr/adb"
)

type Device struct {
	Serial string
}

func ConnectedDevices(includeDeviceSerials []string) ([]Device, error) {
	allDevices, err := AllConnectedDevices()
	if err != nil {
		return nil, err
	}

	if includeDeviceSerials == nil {
		return allDevices, nil
	}

	includedSerials := toMap(includeDeviceSerials)

	var filteredDevices []Device
	for _, device := range allDevices {
		if includedSerials[device.Serial] {
			filteredDevices = append(filteredDevices, device)
		}
	}

	return filteredDevices, nil
}

func FromSerials(serials []string) []Device {
	var devices []Device
	for _, serial := range serials {
		d := Device{
			Serial: serial,
		}
		devices = append(devices, d)
	}
	return devices
}

func AllConnectedDevices() ([]Device, error) {
	deviceSerials, err := adb.ConnectedDevices()
	if err != nil {
		return nil, err
	}

	return FromSerials(deviceSerials), nil
}

func toMap(keys []string) map[string]bool {
	result := map[string]bool{}
	for _, key := range keys {
		result[key] = true
	}
	return result
}
