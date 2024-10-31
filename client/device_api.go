// client/device_api.go

package client

import (
	"context"
	"fmt"

	"github.com/ThingsPanel/tp-protocol-sdk-go/types"
)

// DeviceAPI 设备相关API封装
type DeviceAPI struct {
	client *APIClient
}

// DeviceConfigRequest 获取设备配置请求
type DeviceConfigRequest struct {
	DeviceID     string `json:"device_id"`
	Voucher      string `json:"voucher"`
	DeviceNumber string `json:"device_number"`
}

// DeviceConfigResponse 设备配置响应
type DeviceConfigResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    types.Device `json:"data"`
}

// NewDeviceAPI 创建设备API实例
func NewDeviceAPI(client *APIClient) *DeviceAPI {
	return &DeviceAPI{
		client: client,
	}
}

// GetDeviceConfig 获取设备配置信息
func (d *DeviceAPI) GetDeviceConfig(ctx context.Context, req *DeviceConfigRequest) (*DeviceConfigResponse, error) {
	d.client.logger.Printf("开始获取设备配置: deviceID=%s", req.DeviceID)

	var resp DeviceConfigResponse
	err := d.client.Post(ctx, "/api/v1/plugin/device/config", req, &resp)
	if err != nil {
		d.client.logger.Printf("获取设备配置失败: %v", err)
		return nil, fmt.Errorf("获取设备配置失败: %w", err)
	}

	d.client.logger.Printf("获取设备配置成功: deviceID=%s, deviceType=%s",
		req.DeviceID, resp.Data.DeviceType)
	return &resp, nil
}
