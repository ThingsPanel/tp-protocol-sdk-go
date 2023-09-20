package main

import (
	"fmt"
	"log"
	"time"

	tpprotocolsdkgo "github.com/ThingsPanel/tp-protocol-sdk-go"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// 获取uuid

	// 创建新的MQTT客户端实例
	client := tpprotocolsdkgo.NewMQTTClient("dev.thingspanel.cn:1883", "root", "root")

	// 尝试连接到MQTT代理
	if err := client.Connect(); err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	fmt.Println("连接成功")

	// 订阅一个主题，并提供一个回调函数来处理接收到的消息
	if err := client.Subscribe("test/topic", func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("收到消息: %s\n", msg.Payload())
	}, 0); err != nil {
		log.Fatalf("订阅失败: %v", err)
	}
	// 订阅一个主题，并提供一个回调函数来处理接收到的消息
	if err := client.Subscribe("device/status", func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("收到消息: %s\n", msg.Payload())
	}, 0); err != nil {
		log.Fatalf("订阅失败: %v", err)
	}
	err := client.SendStatus("123", "1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("发送成功")
	// 在一个无限循环中周期性地发布消息
	for {
		if err := client.Publish("test/topic", "Hello, MQTT!", 1); err != nil {
			log.Printf("发布失败: %v", err)
		}
		time.Sleep(1 * time.Second) // 每秒发送一次
	}

}
