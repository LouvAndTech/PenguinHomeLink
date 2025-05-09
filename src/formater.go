package main

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/iancoleman/strcase"
)

// component represents a structure used to define metadata for a specific component.
// It includes fields for identifying the component, its platform, device class,
// unit of measurement, value template, unique identifier, and state topic.
//
// Fields:
// - Name: The name of the component.
// - Platform: The platform associated with the component.
// - DeviceClass: The class of the device, used to categorize the component.
// - UnitOfMeasurement: The unit of measurement for the component's value.
// - ValueTemplate: A template used to format the value of the component.
// - UniqueID: A unique identifier for the component.
// - StateTopic: The MQTT topic where the component's state is published.
type component struct {
	Name              string `json:"name"`
	Platform          string `json:"platform"`
	DeviceClass       string `json:"device_class"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	ValueTemplate     string `json:"value_template"`
	UniqueID          string `json:"unique_id"`
	StateTopic        string `json:"state_topic"`
	Icon              string `json:"icon,omitempty"`
}

// autoDiscoveryDeviceMQTT represents the structure for an MQTT auto-discovery device.
// It contains information about the device, its origin, components, and MQTT-specific details.
//
// Fields:
//
//	Device:
//	  - Identifiers: A list of unique identifiers for the device.
//	  - Name: The name of the device.
//	  - Manufacturer: The manufacturer of the device.
//	  - Model: The model of the device.
//	  - SerialNumber: The serial number of the device.
//
//	Origin:
//	  - Name: The name of the origin source.
//	  - Sw: The software version of the origin source.
//	  - Url: The URL of the origin source.
//
//	Components:
//	  - A map of component names to their respective component details.
//
//	StateTopic:
//	  - The MQTT topic where the device's state is published.
//
//	QoS:
//	  - The Quality of Service (QoS) level for MQTT communication.
type autoDiscoveryDeviceMQTT struct {
	Device struct {
		Identifiers  []string `json:"identifiers"`
		Name         string   `json:"name"`
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		SerialNumber string   `json:"serial_number"`
	} `json:"device"`
	Origin struct {
		Name string `json:"name"`
		Sw   string `json:"sw"`
		Url  string `json:"url"`
	} `json:"o"`
	Components map[string]component `json:"cmps"`
	StateTopic string               `json:"state_topic"`
	QoS        int                  `json:"qos"`
}

// FormatMQTTConfig formats the MQTT configuration for a given device into a JSON string.
// It creates an auto-discovery MQTT structure containing device information, origin details,
// and sensor components.
//
// Parameters:
//   - device: A pointer to a Device object containing the device and sensor information.
//
// Returns:
//   - A JSON string representing the MQTT configuration.
//   - An error if the JSON marshalling fails or any other issue occurs.
func FormatMQTTConfig(device *Device) (string, error) {
	// Create the auto discovery device MQTT structure
	autoDiscoveryDevice := autoDiscoveryDeviceMQTT{
		StateTopic: device.GetDeviceInfo().SerialNumber,
		QoS:        1,
	}

	// Fill the device information
	autoDiscoveryDevice.Device.Identifiers = []string{device.GetDeviceInfo().SerialNumber}
	autoDiscoveryDevice.Device.Name = device.GetDeviceInfo().Name
	autoDiscoveryDevice.Device.Manufacturer = device.GetDeviceInfo().Manufacturer
	autoDiscoveryDevice.Device.Model = device.GetDeviceInfo().Model
	autoDiscoveryDevice.Device.SerialNumber = device.GetDeviceInfo().SerialNumber

	// Fill the origin information
	autoDiscoveryDevice.Origin.Name = SOFTWARE_NAME
	autoDiscoveryDevice.Origin.Sw = SOFTWARE_VERSION
	autoDiscoveryDevice.Origin.Url = SOFTWARE_URL

	// Fill the components information

	autoDiscoveryDevice.Components = make(map[string]component)
	for _, sensor := range device.GetSensors() {
		component := component{
			Name:              sensor.config.Name,
			Platform:          "sensor",
			DeviceClass:       sensor.config.DeviceClass,
			UnitOfMeasurement: sensor.config.UnitOfMeasurement,
			ValueTemplate:     "{{ value_json." + strcase.ToSnake(sensor.config.Name) + " }}",
			UniqueID:          sensor.config.Name + "_" + device.GetDeviceInfo().Name,
			StateTopic:        GetStateTopic(device),
			Icon:              sensor.config.Icon,
		}
		autoDiscoveryDevice.Components[strcase.ToSnake(sensor.config.Name)] = component
	}

	jsonData, err := json.Marshal(autoDiscoveryDevice)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// FormatMQTTValues formats the sensor values of a given device into a JSON string.
// Each sensor value is rounded to two decimal places and converted to a snake_case key.
//
// Parameters:
//   - device: A pointer to the Device object containing the sensors.
//
// Returns:
//   - A JSON string representation of the sensor values with snake_case keys.
//   - An error if any issue occurs during value retrieval, conversion, or JSON marshaling.
func FormatMQTTValues(device *Device) (string, error) {
	// Create the state values
	stateValues := map[string]float64{}

	// Fill the state values
	for _, sensor := range device.GetSensors() {
		value, err := sensor.GetSensorValue()
		if err != nil {
			return "", err
		}
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return "", err
		}
		floatValue = math.Round(floatValue*100) / 100
		stateValues[strcase.ToSnake(sensor.config.Name)] = floatValue
	}

	jsonData, err := json.Marshal(stateValues)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// GetConfigTopic generates the configuration topic string for a given device.
// The topic is constructed using the device's serial number and a predefined
// software name constant.
//
// Parameters:
//   - device: A pointer to the Device object containing device-specific information.
//
// Returns:
//
//	A string representing the configuration topic for the device.
func GetConfigTopic(device *Device) string {
	return "homeassistant/device/" + SOFTWARE_NAME + "/" + device.GetDeviceInfo().SerialNumber + "/config"
}

// GetStateTopic generates the MQTT state topic for a given device.
// The topic is constructed using the software name, the device's serial number,
// and the "state" suffix.
//
// Parameters:
//   - device: A pointer to the Device object for which the state topic is generated.
//
// Returns:
//
//	A string representing the MQTT state topic for the specified device.
func GetStateTopic(device *Device) string {
	return SOFTWARE_NAME + "/" + device.GetDeviceInfo().SerialNumber + "/state"
}
