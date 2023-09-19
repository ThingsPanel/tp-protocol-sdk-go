package main

import (
	"fmt"
	"log"

	tpprotocolsdkgo "github.com/ThingsPanel/tp-protocol-sdk-go"
	"github.com/ThingsPanel/tp-protocol-sdk-go/api"
)

func main() {
	// 创建一个新的客户端实例
	client := tpprotocolsdkgo.NewClient("http://dev.thingspanel.cn") // 替换为你的baseURL

	// 创建一个新的设备配置请求
	request := api.DeviceConfigRequest{
		DeviceID:    "",                                     // 替换为你的DeviceID
		AccessToken: "014f5e6c-e96e-dcaf-f183-0d047464442c", // 替换为你的AccessToken
	}

	// 使用客户端来获取设备配置
	responseData, err := client.API.GetDeviceConfig(request)
	if err != nil {
		log.Fatalf("failed to get device config: %v", err)
	}

	// 打印设备配置响应数据
	fmt.Printf("Device Config Response Data: %+v\n", responseData)

}
