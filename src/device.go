package main

// deviceConfig represents the configuration details of a device.
// It includes the device's name, manufacturer, model, and serial number.
type deviceConfig struct {
	Name         string
	Manufacturer string
	Model        string
	SerialNumber string
}

// Device represents a physical or virtual device in the system.
// It contains configuration details and a collection of associated sensors.
type Device struct {
	config  *deviceConfig
	sensors []*Sensor
}

// NewDevice creates and returns a new instance of a Device with the specified
// name, manufacturer, model, and serial number. The Device is initialized with
// an empty list of sensors.
//
// Parameters:
//   - name: The name of the device.
//   - manufacturer: The manufacturer of the device.
//   - model: The model of the device.
//   - sn: The serial number of the device.
//
// Returns:
//
//	A pointer to the newly created Device instance.
func NewDevice(name string, manufacturer string, model string, sn string) *Device {
	return &Device{
		config: &deviceConfig{
			Name:         name,
			Manufacturer: manufacturer,
			Model:        model,
			SerialNumber: sn,
		},
		sensors: []*Sensor{},
	}
}

// AddSensor adds a new sensor to the device's list of sensors.
// It appends the provided sensor to the internal slice of sensors.
//
// Parameters:
//   - sensor: A pointer to the Sensor object to be added.
func (d *Device) AddSensor(sensor *Sensor) {
	d.sensors = append(d.sensors, sensor)
}

// GetDeviceInfo retrieves the configuration information of the device.
// It returns a pointer to the deviceConfig struct associated with the device.
func (d *Device) GetDeviceInfo() *deviceConfig {
	return d.config
}

// GetSensors returns a slice of pointers to Sensor objects associated with the Device.
// This method provides access to the sensors managed by the Device instance.
func (d *Device) GetSensors() []*Sensor {
	return d.sensors
}
