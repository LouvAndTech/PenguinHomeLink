package main

import (
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
)

type mqttConfig struct {
	IP       string
	Port     string
	Username string
	Password string
}

type MQTTProxy struct {
	IsConnected bool
	config      *mqttConfig
	opts        *mqtt.ClientOptions
	client      mqtt.Client
}

func NewMQTTProxy(ip string, port string, username string, password string) *MQTTProxy {
	return &MQTTProxy{
		IsConnected: false,
		config: &mqttConfig{
			IP:       ip,
			Port:     port,
			Username: username,
			Password: password,
		},
		opts:   nil,
		client: nil,
	}
}

func (m *MQTTProxy) Connect() error {
	if m.IsConnected {
		return nil
	}
	m.opts = mqtt.NewClientOptions()
	m.opts.AddBroker("tcp://" + m.config.IP + ":" + m.config.Port)
	m.opts.SetUsername(m.config.Username)
	m.opts.SetPassword(m.config.Password)
	m.opts.AutoReconnect = true

	// Set the OnConnectionLost callback
	m.opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
		m.IsConnected = false
	})

	m.client = mqtt.NewClient(m.opts)

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	m.IsConnected = true
	return nil
}

func (m *MQTTProxy) Disconnect() {
	if m.IsConnected {
		m.client.Disconnect(250)
		m.IsConnected = false
	}
}

func (m *MQTTProxy) Publish(topic string, payload string) error {
	if !m.IsConnected {
		return fmt.Errorf("not connected to MQTT broker")
	}

	token := m.client.Publish(topic, 0, false, payload)
	token.Wait()

	if token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *MQTTProxy) Subscribe(topic string, callback mqtt.MessageHandler) error {
	if !m.IsConnected {
		return fmt.Errorf("not connected to MQTT broker")
	}

	token := m.client.Subscribe(topic, 0, callback)
	token.Wait()

	if token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *MQTTProxy) Unsubscribe(topic string) error {
	if !m.IsConnected {
		return fmt.Errorf("not connected to MQTT broker")
	}

	token := m.client.Unsubscribe(topic)
	token.Wait()

	if token.Error() != nil {
		return token.Error()
	}

	return nil
}
