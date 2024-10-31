// examples/main.go

package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ThingsPanel/tp-protocol-sdk-go/client"
)

func main() {
	// 创建logger
	logger := log.New(os.Stdout, "[TP-Example] ", log.LstdFlags|log.Lshortfile)

	// 创建客户端配置
	config := client.ClientConfig{
		BaseURL:      "http://c.thingspanel.cn",
		MQTTBroker:   "mqtt://47.92.253.145:1883",
		MQTTClientID: "example-client",
		MQTTUsername: "plugin",
		MQTTPassword: "plugin",
		Logger:       logger,
	}

	// 创建SDK客户端
	c, err := client.NewClient(config)
	if err != nil {
		logger.Fatalf("创建客户端失败: %v", err)
	}
	defer c.Close()

	// 连接平台
	if err := c.Connect(); err != nil {
		logger.Fatalf("连接平台失败: %v", err)
	}

	// 获取设备配置示例
	deviceReq := &client.DeviceConfigRequest{
		DeviceID: "af13ac2c-3a9e-5ab9-cd31-0cf01f984b3c",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deviceResp, err := c.Device().GetDeviceConfig(ctx, deviceReq)
	if err != nil {
		logger.Printf("获取设备配置失败: %v", err)
	} else {
		logger.Printf("设备配置: %+v", deviceResp.Data)
	}

	// 发送服务心跳示例
	heartbeatReq := &client.HeartbeatRequest{
		ServiceIdentifier: "test3",
	}

	heartbeatResp, err := c.Service().SendHeartbeat(ctx, heartbeatReq)
	if err != nil {
		logger.Printf("发送心跳失败: %v", err)
	} else {
		logger.Printf("心跳响应: %+v", heartbeatResp)
	}

	// MQTT消息订阅示例
	err = c.MQTT().Subscribe("plugin/example/+", 1, func(topic string, payload []byte) {
		logger.Printf("收到消息: topic=%s, payload=%s", topic, string(payload))
	})
	if err != nil {
		logger.Printf("订阅主题失败: %v", err)
	}

	// 发布消息示例
	err = c.MQTT().Publish("devices/status/af13ac2c-3a9e-5ab9-cd31-0cf01f984b3c", 1, 1)
	if err != nil {
		logger.Printf("发布消息失败: %v", err)
	}

	// 保持运行一段时间
	time.Sleep(600 * time.Second)
}
