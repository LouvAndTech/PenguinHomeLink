package main

import (
	"fmt"

	"github.com/eclipse/paho.mqtt.golang"
)

// mqttConfig represents the configuration required to connect to an MQTT broker.
// It includes the following fields:
// - IP: The IP address of the MQTT broker.
// - Port: The port number on which the MQTT broker is running.
// - Username: The username for authenticating with the MQTT broker.
// - Password: The password for authenticating with the MQTT broker.
type mqttConfig struct {
	IP       string
	Port     string
	Username string
	Password string
}

// MQTTProxy represents a proxy for managing MQTT client connections and configurations.
// It provides functionality to maintain the connection state, client options, and the MQTT client instance.
//
// Fields:
// - IsConnected: A boolean indicating whether the MQTT client is currently connected.
// - config: A private field containing the configuration details for the MQTT client.
// - opts: A private field holding the MQTT client options.
// - client: A private field representing the MQTT client instance.
type MQTTProxy struct {
	IsConnected bool
	config      *mqttConfig
	opts        *mqtt.ClientOptions
	client      mqtt.Client
}

// NewMQTTProxy creates a new instance of MQTTProxy with the specified configuration.
// It initializes the MQTTProxy with the provided IP address, port, username, and password.
//
// Parameters:
//   - ip: The IP address of the MQTT broker.
//   - port: The port number of the MQTT broker.
//   - username: The username for authenticating with the MQTT broker.
//   - password: The password for authenticating with the MQTT broker.
//
// Returns:
//
//	A pointer to an initialized MQTTProxy instance.
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

// Connect establishes a connection to the MQTT broker using the configuration
// provided in the MQTTProxy instance. If the connection is already established,
// the method returns immediately without performing any action.
//
// The method sets up client options, including the broker address, username,
// password, and auto-reconnect behavior. It also defines a callback to handle
// connection loss, which updates the connection status and logs the error.
//
// Returns an error if the connection attempt fails; otherwise, it updates the
// IsConnected field to true upon successful connection.
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

// Disconnect gracefully disconnects the MQTT client if it is currently connected.
// It waits for up to 250 milliseconds to ensure any pending operations are completed
// before closing the connection. After disconnecting, it updates the IsConnected
// flag to false.
func (m *MQTTProxy) Disconnect() {
	if m.IsConnected {
		m.client.Disconnect(250)
		m.IsConnected = false
	}
}

// Publish sends a message with the specified payload to the given MQTT topic.
// It ensures that the MQTT client is connected before attempting to publish.
// If the client is not connected, or if an error occurs during publishing,
// an error is returned.
//
// Parameters:
//   - topic: The MQTT topic to which the message will be published.
//   - payload: The message content to be published.
//
// Returns:
//   - error: An error if the client is not connected or if the publish operation fails.
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

// Subscribe subscribes to a specific MQTT topic and registers a callback function
// to handle incoming messages on that topic.
//
// Parameters:
//   - topic: The MQTT topic to subscribe to.
//   - callback: A function of type mqtt.MessageHandler that will be invoked
//     whenever a message is received on the subscribed topic.
//
// Returns:
//   - error: An error if the subscription fails, or if the MQTT client is not connected.
//
// Notes:
//   - The function checks if the MQTT client is connected before attempting to subscribe.
//   - The callback function is executed for each message received on the subscribed topic.
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

// Unsubscribe unsubscribes the MQTT client from the specified topic.
// It ensures that the client is connected before attempting to unsubscribe.
// If the client is not connected, it returns an error.
// If the unsubscription process encounters an error, it returns the error.
//
// Parameters:
//   - topic: The MQTT topic to unsubscribe from.
//
// Returns:
//   - error: An error if the client is not connected or if the unsubscription fails; otherwise, nil.
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
