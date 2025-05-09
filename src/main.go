package main

import (
	"fmt"
	"time"
)

const (
	SOFTWARE_NAME    = "PenguinHomeLink"
	SOFTWARE_VERSION = "1.0.0"
	SOFTWARE_URL     = "https://github.com/LouvAndTech/PenguinHomeLink"

	RETRY_PAUSE           = 30 * time.Second // 2 * time.Minute
	CONFIG_REFRESH_PERIOD = 15 * time.Minute
)

func main() {
	fmt.Println("Starting PenguinHomeLink...")

	fmt.Println(">> Configuring...")
	// Parse the confuguration file
	fmt.Println(">> Loading configuration...")
	config, err := LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	// Print the configuration
	fmt.Println(">> Configuration loaded successfully.")
	//fmt.Printf("%+v\n", config)

	//create the device
	fmt.Println(">> Creating device and sensors...")
	device := NewDevice(config.Device.Name, config.Device.Manufacturer, config.Device.Model, config.Device.SerialNumber)
	// Print the device information
	//fmt.Printf("%+v\n", device.GetDeviceInfo())

	// Create the sensors
	for _, sensorConfig := range config.Sensors {
		sensor := NewSensor(sensorConfig.Name, sensorConfig.Command, sensorConfig.DeviceClass, sensorConfig.StateClass, sensorConfig.UnitOfMeasurement, sensorConfig.Icon, device)
		device.AddSensor(sensor)
	}
	// Print the sensors information
	// for _, sensor := range device.GetSensors() {
	// 	fmt.Printf("%+v\n", sensor.config)
	// }
	fmt.Println(">> Device and sensors created successfully.")

	// create the MQTT server proxy
	fmt.Println(">> Creating MQTT server proxy...")
	MQTTServer := NewMQTTProxy(config.MQTTServer.IP, config.MQTTServer.Port, config.MQTTServer.Username, config.MQTTServer.Password)
	fmt.Println(">> MQTT server proxy created successfully.")

	fmt.Println(">> Software configured successfully.")

	// Run the main loop
	fmt.Println(">> Running...")
	run(device, MQTTServer, config.Software.RefreshPeriodS)
}

func run(device *Device, MQTTServer *MQTTProxy, refreshPeriod int) {

	// Format the MQTT config payload
	mqttConfig, err := FormatMQTTConfig(device)
	if err != nil {
		panic(err)
	}

	lastConfigSent := time.Date(2020, 10, 26, 0, 0, 0, 0, time.UTC)
	for {
		func() {
			//If an error occurs, wait for 2 mins before trying again
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(">>> Recovered in run:", r)
					fmt.Println(">>> Waiting for 2 minutes before retrying...")
					time.Sleep(RETRY_PAUSE)
				}
			}()

			// Connect to the MQTT server
			err := MQTTServer.Connect()
			if err != nil {
				panic(err)
			}
			fmt.Println("> Connected to the MQTT server.")

			// Send configuration to the MQTT server if 15 minutes have elapsed
			if time.Since(lastConfigSent) > CONFIG_REFRESH_PERIOD {
				fmt.Println("> Sending configuration to the MQTT server...")

				err = MQTTServer.Publish(GetConfigTopic(device), mqttConfig)
				if err != nil {
					panic(err)
				}
				fmt.Println("> Configuration sent to the MQTT server.")
				lastConfigSent = time.Now()
			}

			fmt.Println("> Getting sensor values...")
			for _, sensor := range device.GetSensors() {
				val, err := sensor.GetSensorValue()
				if err != nil {
					fmt.Println("Error getting sensor value:", err)
					continue
				}
				fmt.Println("Sensor : ", sensor.config.Name, " - value:", val)
			}
			fmt.Println("> Sensor values retrieved successfully.")

			// Format the MQTT values payload
			fmt.Println("> Sending sensor values to the MQTT server...")
			mqttValues, err := FormatMQTTValues(device)
			if err != nil {
				panic(err)
			}

			// Publish sensor values to the MQTT server
			err = MQTTServer.Publish(GetStateTopic(device), mqttValues)
			if err != nil {
				panic(err)
			}
			fmt.Println("> Sensor values sent to the MQTT server.")

			// Sleep for a defined interval before the next iteration
			time.Sleep(time.Duration(refreshPeriod) * time.Second) // Adjust the interval as needed
		}()
	}
}
