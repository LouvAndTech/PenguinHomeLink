package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the configuration file.
type Config struct {
	Software struct {
		RefreshPeriodS int `yaml:"refresh_period_s"`
	} `yaml:"software"`

	Device struct {
		Name         string `yaml:"name"`
		Manufacturer string `yaml:"manufacturer"`
		Model        string `yaml:"model"`
		SerialNumber string `yaml:"serial_number"`
	} `yaml:"device"`

	MQTTServer struct {
		IP       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mqtt_server"`

	Sensors []struct {
		Name              string `yaml:"name"`
		Command           string `yaml:"command"`
		DeviceClass       string `yaml:"device_class"`
		StateClass        string `yaml:"state_class"`
		UnitOfMeasurement string `yaml:"unit_of_measurement"`
	} `yaml:"sensors"`
}

// LoadConfig reads a YAML file from the given path and unmarshals it into a Config struct.
func LoadConfig(filePath string) (*Config, error) {
	// Open the YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Decode the YAML file
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	return &config, nil
}
