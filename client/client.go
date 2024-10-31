// client/client.go

package client

import (
	"fmt"
	"log"
)

// Client SDK主客户端，整合所有功能
type Client struct {
	// API客户端
	api     *APIClient
	device  *DeviceAPI
	service *ServiceAPI

	// MQTT客户端
	mqtt *MQTTClient

	// 日志
	logger *log.Logger
}

// ClientConfig SDK客户端配置
type ClientConfig struct {
	// API配置
	BaseURL    string
	APITimeout int // 秒

	// MQTT配置
	MQTTBroker   string
	MQTTClientID string
	MQTTUsername string
	MQTTPassword string

	// 日志配置
	Logger *log.Logger
}

// NewClient 创建新的SDK客户端实例
func NewClient(config ClientConfig) (*Client, error) {
	// 设置默认logger
	logger := config.Logger
	if logger == nil {
		logger = log.New(log.Writer(), "[TP-SDK] ", log.LstdFlags|log.Lshortfile)
	}

	logger.Printf("初始化SDK客户端")

	// 创建API客户端
	apiClient := NewAPIClient(config.BaseURL, WithLogger(logger))
	if apiClient == nil {
		return nil, fmt.Errorf("创建API客户端失败")
	}

	// 创建MQTT客户端
	mqttClient := NewMQTTClient(MQTTConfig{
		Broker:   config.MQTTBroker,
		ClientID: config.MQTTClientID,
		Username: config.MQTTUsername,
		Password: config.MQTTPassword,
	}, logger)
	if mqttClient == nil {
		return nil, fmt.Errorf("创建MQTT客户端失败")
	}

	// 创建设备API和服务API
	deviceAPI := NewDeviceAPI(apiClient)
	serviceAPI := NewServiceAPI(apiClient)

	return &Client{
		api:     apiClient,
		device:  deviceAPI,
		service: serviceAPI,
		mqtt:    mqttClient,
		logger:  logger,
	}, nil
}

// Connect 连接到平台
func (c *Client) Connect() error {
	c.logger.Printf("开始连接平台")

	// 连接MQTT
	if err := c.mqtt.Connect(); err != nil {
		c.logger.Printf("MQTT连接失败: %v", err)
		return fmt.Errorf("MQTT连接失败: %w", err)
	}

	c.logger.Printf("平台连接成功")
	return nil
}

// Device 获取设备API操作接口
func (c *Client) Device() *DeviceAPI {
	return c.device
}

// Service 获取服务API操作接口
func (c *Client) Service() *ServiceAPI {
	return c.service
}

// MQTT 获取MQTT客户端
func (c *Client) MQTT() *MQTTClient {
	return c.mqtt
}

// Close 关闭客户端连接
func (c *Client) Close() {
	c.logger.Printf("开始关闭客户端连接")

	// 断开MQTT连接
	if c.mqtt != nil {
		c.mqtt.Disconnect()
	}

	c.logger.Printf("客户端连接已关闭")
}
