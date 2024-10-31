// client/mqtt_client.go

package client

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTTClient MQTT客户端封装
type MQTTClient struct {
	client    mqtt.Client
	logger    *log.Logger
	broker    string
	clientID  string
	username  string
	password  string
	connected bool
}

// MessageHandler 定义消息处理函数类型
type MessageHandler func(topic string, payload []byte)

// MQTTConfig MQTT配置项
type MQTTConfig struct {
	Broker   string
	ClientID string
	Username string
	Password string
}

// NewMQTTClient 创建MQTT客户端实例
func NewMQTTClient(config MQTTConfig, logger *log.Logger) *MQTTClient {
	if logger == nil {
		logger = log.New(log.Writer(), "[TP-MQTT] ", log.LstdFlags|log.Lshortfile)
	}

	return &MQTTClient{
		broker:   config.Broker,
		clientID: config.ClientID,
		username: config.Username,
		password: config.Password,
		logger:   logger,
	}
}

// Connect 连接到MQTT服务器
func (m *MQTTClient) Connect() error {
	m.logger.Printf("开始连接MQTT服务器: broker=%s, clientID=%s", m.broker, m.clientID)

	opts := mqtt.NewClientOptions().
		AddBroker(m.broker).
		SetClientID(m.clientID).
		SetUsername(m.username).
		SetPassword(m.password).
		SetAutoReconnect(true).
		SetCleanSession(true).
		SetKeepAlive(30 * time.Second).
		SetConnectTimeout(30 * time.Second)

	// 设置连接丢失处理函数
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		m.logger.Printf("MQTT连接丢失: %v", err)
		m.connected = false
	})

	// 设置连接建立处理函数
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		m.logger.Printf("MQTT连接成功建立")
		m.connected = true
	})

	// 创建客户端实例
	m.client = mqtt.NewClient(opts)

	// 尝试连接
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		m.logger.Printf("MQTT连接失败: %v", token.Error())
		return fmt.Errorf("MQTT连接失败: %w", token.Error())
	}

	m.logger.Printf("MQTT连接成功")
	return nil
}

// Publish 发布消息
func (m *MQTTClient) Publish(topic string, qos byte, payload interface{}) error {
	if !m.connected {
		return fmt.Errorf("MQTT客户端未连接")
	}

	m.logger.Printf("准备发布消息: topic=%s, qos=%d", topic, qos)

	token := m.client.Publish(topic, qos, false, payload)
	if token.Wait() && token.Error() != nil {
		m.logger.Printf("消息发布失败: %v", token.Error())
		return fmt.Errorf("消息发布失败: %w", token.Error())
	}

	m.logger.Printf("消息发布成功")
	return nil
}

// Subscribe 订阅主题
func (m *MQTTClient) Subscribe(topic string, qos byte, handler MessageHandler) error {
	if !m.connected {
		return fmt.Errorf("MQTT客户端未连接")
	}

	m.logger.Printf("准备订阅主题: topic=%s, qos=%d", topic, qos)

	// 将自定义的MessageHandler转换为mqtt.MessageHandler
	wrapper := func(client mqtt.Client, msg mqtt.Message) {
		handler(msg.Topic(), msg.Payload())
	}

	token := m.client.Subscribe(topic, qos, wrapper)
	if token.Wait() && token.Error() != nil {
		m.logger.Printf("主题订阅失败: %v", token.Error())
		return fmt.Errorf("主题订阅失败: %w", token.Error())
	}

	m.logger.Printf("主题订阅成功")
	return nil
}

// Disconnect 断开MQTT连接
func (m *MQTTClient) Disconnect() {
	if m.connected {
		m.logger.Printf("准备断开MQTT连接")
		m.client.Disconnect(250)
		m.connected = false
		m.logger.Printf("MQTT连接已断开")
	}
}

// IsConnected 检查是否已连接
func (m *MQTTClient) IsConnected() bool {
	return m.connected
}
