package devicebattery

type DeviceTypes struct {
	AllDevices map[string]DeviceType
}

type DeviceType struct {
	DeviceTitles []map[string]DeviceSpecs
}

// DeviceSpecs contains attributes of individual devices
type DeviceSpecs struct {
	BatteryPercent string `json:"device_batteryPercent,omitempty"`
	Services       string `json:"device_services,omitempty"`
}
