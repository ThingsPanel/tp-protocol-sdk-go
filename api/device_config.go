package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 请求结构体
type DeviceConfigRequest struct {
	DeviceID     string `json:"device_id"`
	Voucher      string `json:"voucher"`
	DeviceNumber string `json:"device_number"`
}

// 子设备信息结构体
type SubDevice struct {
	Voucher                string                 `json:"voucher"`
	DeviceID               string                 `json:"device_id"`
	SubDeviceAddr          string                 `json:"sub_device_addr"`
	Config                 map[string]interface{} `json:"config"`
	ProtocolConfigTemplate map[string]interface{} `json:"protocol_config_template"` // 子设备配置的protocol_config
}

// 设备信息结构体
type DeviceConfigResponseData struct {
	Voucher                string                 `json:"voucher"`
	DeviceType             string                 `json:"device_type"`
	ID                     string                 `json:"id"`
	ProtocolType           string                 `json:"Protocol_type"`
	SubDevices             []SubDevice            `json:"sub_devices"`
	Config                 map[string]interface{} `json:"config"`
	ProtocolConfigTemplate map[string]interface{} `json:"protocol_config_template"` // 子设备配置的protocol_config
}

// 响应结构体
type DeviceConfigResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    DeviceConfigResponseData `json:"data"`
}

// 获取设备配置（通过设备id或者设备令牌）
func (a *API) GetDeviceConfig(request DeviceConfigRequest) (*DeviceConfigResponse, error) {

	// 构建API终端的URL
	apiEndpoint := fmt.Sprintf("%s/api/v1/plugin/device/config", a.BaseURL)

	// 向API发送POST请求
	resp, err := a.doPostRequest(apiEndpoint, request)
	if err != nil {
		return nil, err
	}
	// 确保在处理完响应后关闭响应体。
	defer resp.Body.Close()

	// 检查HTTP状态码，确保我们收到了成功的响应。
	// 如果不是，返回一个包含状态码的错误信息。
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("获取设备配置失败: HTTP状态码 %d", resp.StatusCode)
	}

	// 读取整个响应体。
	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 初始化一个变量来存放响应数据结构。
	var response DeviceConfigResponse

	// 解析响应中的JSON数据，并填充到我们的结构体中。
	err = json.Unmarshal(p, &response)
	if err != nil {
		return nil, fmt.Errorf("解析响应JSON失败: %w; 响应内容: %s", err, string(p))
	}

	// 返回填充后的响应结构体。
	return &response, nil
}
