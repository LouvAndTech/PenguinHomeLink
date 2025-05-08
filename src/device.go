package main

type deviceConfig struct {
	Name         string
	Manufacturer string
	Model        string
	SerialNumber string
}

type Device struct {
	config  *deviceConfig
	sensors []*Sensor
}

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

func (d *Device) AddSensor(sensor *Sensor) {
	d.sensors = append(d.sensors, sensor)
}

func (d *Device) GetDeviceInfo() *deviceConfig {
	return d.config
}

func (d *Device) GetSensors() []*Sensor {
	return d.sensors
}
