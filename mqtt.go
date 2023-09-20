package tpprotocolsdkgo

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-basic/uuid"
)

// MQTT客户端结构体
type MQTTClient struct {
	client mqtt.Client
}

// 创建新的MQTT客户端
func NewMQTTClient(broker string, username string, password string) *MQTTClient {
	clientID := uuid.New()
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true) // 启用自动重新连接
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("连接丢失: %v", err)
	})

	client := mqtt.NewClient(opts)
	return &MQTTClient{client: client}
}

// 连接到MQTT代理，如果连接失败则重试100次
func (client *MQTTClient) Connect() error {
	for retry := 0; retry < 100; retry++ {
		if token := client.client.Connect(); token.Wait() && token.Error() == nil {
			return nil
		}
		log.Printf("连接失败, 尝试重新连接, 尝试 #%d...", retry+1)
		time.Sleep(6 * time.Second) // 等待6秒后重试
	}
	return fmt.Errorf("连接失败: 达到最大重试次数")
}

// 发布消息到指定主题
func (client *MQTTClient) Publish(topic string, payload string, qos uint8) error {
	token := client.client.Publish(topic, qos, false, payload)
	token.Wait()
	return token.Error()
}

// 订阅指定主题，并提供一个处理接收到消息的回调函数
func (client *MQTTClient) Subscribe(topic string, callback mqtt.MessageHandler, qos uint8) error {
	if token := client.client.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

//发送在线离线消息status 1-在线 0-离线
func (client *MQTTClient) SendStatus(accessToken string, status string) (err error) {
	// 校验参数
	if accessToken == "" {
		return fmt.Errorf("accessToken不能为空")
	}
	if status != "1" && status != "0" {
		return fmt.Errorf("status只能为1或0")
	}
	payload := `{"accessToken":"` + accessToken + `","values":{"status":"` + status + `"}}`
	token := client.client.Publish("device/status", 1, false, string(payload))
	token.Wait()
	return token.Error()
}
