package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the configuration structure for the application.
// It includes settings for software, device information, MQTT server, and sensors.
//
// Fields:
// - Software: Contains software-related configurations such as the refresh period.
//   - RefreshPeriodS: The refresh period in seconds.
//
// - Device: Contains information about the device.
//   - Name: The name of the device.
//   - Manufacturer: The manufacturer of the device.
//   - Model: The model of the device.
//   - SerialNumber: The serial number of the device.
//
// - MQTTServer: Contains the configuration for the MQTT server.
//   - IP: The IP address of the MQTT server.
//   - Port: The port of the MQTT server.
//   - Username: The username for MQTT server authentication.
//   - Password: The password for MQTT server authentication.
//
// - Sensors: A list of sensor configurations.
//   - Name: The name of the sensor.
//   - Command: The command associated with the sensor.
//   - DeviceClass: The device class of the sensor.
//   - StateClass: The state class of the sensor.
//   - UnitOfMeasurement: The unit of measurement for the sensor's data.
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
		Icon              string `yaml:"icon,omitempty"`
	} `yaml:"sensors"`
}

// LoadConfig loads the configuration from a YAML file located at the specified file path.
// It opens the file, decodes its contents into a Config struct, and returns a pointer to the struct.
// If an error occurs during file opening or decoding, it returns an error.
//
// Parameters:
//   - filePath: The path to the YAML configuration file.
//
// Returns:
//   - *Config: A pointer to the loaded configuration struct.
//   - error: An error if the file cannot be opened or the contents cannot be decoded.
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
