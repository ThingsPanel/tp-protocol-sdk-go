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
	request := api.DeviceConfigListRequest{
		ProtocolType: "MODBUS_RTU", // 替换为你的ProtocolType
		DeviceType:   "2",          // 替换为你的DeviceType
	}

	// 使用客户端来获取设备配置
	responseData, err := client.API.GetDeviceConfigList(request)
	if err != nil {
		log.Fatalf("failed to get device config: %v", err)
	}

	// 打印设备配置响应数据
	fmt.Printf("Device Config Response Data: %+v\n", responseData)

}
