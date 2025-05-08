package main

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/iancoleman/strcase"
)

type component struct {
	Name              string `json:"name"`
	Platform          string `json:"platform"`
	DeviceClass       string `json:"device_class"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	ValueTemplate     string `json:"value_template"`
	UniqueID          string `json:"unique_id"`
	StateTopic        string `json:"state_topic"`
}

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
		}
		autoDiscoveryDevice.Components[strcase.ToSnake(sensor.config.Name)] = component
	}

	jsonData, err := json.Marshal(autoDiscoveryDevice)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

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

func GetConfigTopic(device *Device) string {
	return "homeassistant/device/" + SOFTWARE_NAME + "/" + device.GetDeviceInfo().SerialNumber + "/config"
}

func GetStateTopic(device *Device) string {
	return SOFTWARE_NAME + "/" + device.GetDeviceInfo().SerialNumber + "/state"
}
