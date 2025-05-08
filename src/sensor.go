package main

import (
	"os/exec"
	"strings"
)

// sensorConfig represents the configuration for a sensor.
// It includes details such as the sensor's name, the command to retrieve its data,
// its device class, state class, and the unit of measurement used.
//
// Fields:
// - Name: The name of the sensor.
// - Command: The command used to retrieve data from the sensor.
// - DeviceClass: The type or category of the sensor (e.g., temperature, humidity).
// - StateClass: The classification of the sensor's state (e.g., measurement, total).
// - UnitOfMeasurement: The unit in which the sensor's data is measured (e.g., °C, %, m/s).
type sensorConfig struct {
	Name              string
	Command           string
	DeviceClass       string
	StateClass        string
	UnitOfMeasurement string
	Icon              string
}

// Sensor represents a sensor device in the system.
// It contains configuration details, the current value of the sensor,
// and a reference to the associated Device.
type Sensor struct {
	config *sensorConfig
	value  string
	Device *Device
}

// NewSensor creates and returns a new Sensor instance with the specified configuration.
//
// Parameters:
//   - name: The name of the sensor.
//   - command: The command associated with the sensor.
//   - deviceClass: The class of the device (e.g., temperature, humidity).
//   - stateClass: The state class of the sensor (e.g., measurement, total).
//   - unitOfMeasurement: The unit of measurement for the sensor's value (e.g., °C, %, etc.).
//   - device: A pointer to the associated Device instance.
//
// Returns:
//   - A pointer to the newly created Sensor instance.
func NewSensor(name string, command string, deviceClass string, stateClass string, unitOfMeasurement string, icon string, device *Device) *Sensor {
	return &Sensor{
		config: &sensorConfig{
			Name:              name,
			Command:           command,
			DeviceClass:       deviceClass,
			StateClass:        stateClass,
			UnitOfMeasurement: unitOfMeasurement,
			Icon:              icon,
		},
		value:  "",
		Device: device,
	}
}

// GetSensorValue retrieves the current value of the sensor after performing a measurement.
// It returns the sensor value as a string if the measurement is successful, or an error
// if the measurement fails.
func (s *Sensor) GetSensorValue() (string, error) {
	err := s.runMeasurement()
	if err != nil {
		return "", err
	}
	return s.value, nil
}

func (s *Sensor) runMeasurement() error {
	cmd := exec.Command("bash", "-c", s.config.Command)
	output, err := cmd.Output()
	if err != nil {
		s.value = ""
		return err
	}
	s.value = strings.TrimSpace(string(output))
	return nil
}
