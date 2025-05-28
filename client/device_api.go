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

// DeviceDynamicAuthRequest 设备动态认证请求
// 对应接口 /api/v1/device/auth
// Body 参数 application/json
// - template_secret (string, 必需)
// - device_number (string, 必需)
// - device_name (string, 可选)
// - product_key (string, 可选)
type DeviceDynamicAuthRequest struct {
	TemplateSecret string `json:"template_secret"`
	DeviceNumber   string `json:"device_number"`
	DeviceName     string `json:"device_name,omitempty"`
	ProductKey     string `json:"product_key,omitempty"`
}

// DeviceDynamicAuthResponse 设备动态认证响应
// 响应示例：
//
//	{
//	  "code": 200,
//	  "message": "操作成功",
//	  "data": {
//	    "device_id": "4e7e16dc-b1e7-5eef-32a6-48a7d767c85f",
//	    "voucher": "{\"username\":\"6c2f1bdc-6fc2-b535-f0ba-f77fe9dc6db1\"}"
//	  }
//	}
type DeviceDynamicAuthResponse struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    types.DeviceDynamicAuthData `json:"data"`
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

// DeviceDynamicAuth 设备动态认证接口
func (d *DeviceAPI) DeviceDynamicAuth(ctx context.Context, req *DeviceDynamicAuthRequest) (*DeviceDynamicAuthResponse, error) {
	d.client.logger.Printf("开始设备动态认证: deviceNumber=%s", req.DeviceNumber)

	var resp DeviceDynamicAuthResponse
	err := d.client.Post(ctx, "/api/v1/device/auth", req, &resp)
	if err != nil {
		d.client.logger.Printf("设备动态认证失败: %v", err)
		return nil, fmt.Errorf("设备动态认证失败: %w", err)
	}
	d.client.logger.Printf("设备动态认证成功: deviceID=%s", resp.Data.DeviceID)
	return &resp, nil
}
