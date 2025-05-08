package main

import (
	"os/exec"
	"strings"
)

type sensorConfig struct {
	Name              string
	Command           string
	DeviceClass       string
	StateClass        string
	UnitOfMeasurement string
}

type Sensor struct {
	config *sensorConfig
	value  string
	Device *Device
}

func NewSensor(name string, command string, deviceClass string, stateClass string, unitOfMeasurement string, device *Device) *Sensor {
	return &Sensor{
		config: &sensorConfig{
			Name:              name,
			Command:           command,
			DeviceClass:       deviceClass,
			StateClass:        stateClass,
			UnitOfMeasurement: unitOfMeasurement,
		},
		value:  "",
		Device: device,
	}
}

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
