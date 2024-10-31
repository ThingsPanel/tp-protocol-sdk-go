// examples/handler/main.go

package main

import (
	"log"
	"os"

	"github.com/ThingsPanel/tp-protocol-sdk-go/handler"
)

func main() {
	logger := log.New(os.Stdout, "[TP-Example] ", log.LstdFlags|log.Lshortfile)

	// 创建处理器
	h := handler.NewHandler(handler.HandlerConfig{
		Logger: logger,
	})

	// 设置表单配置处理函数
	h.SetFormConfigHandler(func(req *handler.GetFormConfigRequest) (interface{}, error) {
		logger.Printf("收到表单配置请求: type=%s", req.FormType)
		return map[string]interface{}{
			"fields": []map[string]interface{}{
				{
					"name":  "host",
					"type":  "string",
					"label": "服务器地址",
				},
				{
					"name":  "port",
					"type":  "number",
					"label": "端口",
				},
			},
		}, nil
	})

	// 设置设备断开连接处理函数
	h.SetDeviceDisconnectHandler(func(req *handler.DeviceDisconnectRequest) error {
		logger.Printf("设备断开连接: %s", req.DeviceID)
		return nil
	})

	// 设置通知处理函数
	h.SetNotificationHandler(func(req *handler.NotificationRequest) error {
		logger.Printf("收到通知: type=%s", req.MessageType)
		return nil
	})

	// 设置获取设备列表处理函数
	h.SetGetDeviceListHandler(func(req *handler.GetDeviceListRequest) (*handler.DeviceListResponse, error) {
		logger.Printf("获取设备列表: page=%d, pageSize=%d", req.Page, req.PageSize)
		return &handler.DeviceListResponse{
			Code:    200,
			Message: "success",
			Data: handler.DeviceListData{
				List: []handler.DeviceItem{
					{
						DeviceName:   "设备1",
						Description:  "测试设备1",
						DeviceNumber: "DEV001",
					},
					{
						DeviceName:   "设备2",
						Description:  "测试设备2",
						DeviceNumber: "DEV002",
					},
				},
				Total: 2,
			},
		}, nil
	})

	// 启动HTTP服务
	logger.Printf("启动HTTP服务...")
	if err := h.Start(":8080"); err != nil {
		logger.Fatalf("服务启动失败: %v", err)
	}
}
