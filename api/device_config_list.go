package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 请求结构体
type DeviceConfigListRequest struct {
	ProtocolType string `json:"protocol_type"`
	DeviceType   string `json:"device_type"`
}

// 子设备信息结构体
type SubDeviceConfig struct {
	AccessToken   string                 `json:"AccessToken"`
	DeviceID      string                 `json:"DeviceId"`
	SubDeviceAddr string                 `json:"SubDeviceAddr"`
	Config        map[string]interface{} `json:"Config"` // 表单配置
}

// 设备信息结构体
type DeviceConfigListResponseData struct {
	ProtocolType string                 `json:"ProtocolType"`
	AccessToken  string                 `json:"AccessToken"`
	DeviceType   string                 `json:"DeviceType"`
	ID           string                 `json:"Id"`
	DeviceConfig map[string]interface{} `json:"DeviceConfig,omitempty"` // 表单配置
	SubDevices   []SubDeviceConfig      `json:"SubDevices,omitempty"`
}

// 响应结构体
type DeviceConfigListResponse struct {
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Data    []DeviceConfigResponseData `json:"data"`
}

// 响应内容会被解析到 DeviceConfigListResponse 结构体中。
func (a *API) GetDeviceConfigList(request DeviceConfigListRequest) (*DeviceConfigListResponse, error) {

	// 构建API终端的URL
	apiEndpoint := fmt.Sprintf("%s/api/plugin/all_device/config", a.BaseURL)

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
		return nil, fmt.Errorf("获取设备配置列表失败: HTTP状态码 %d", resp.StatusCode)
	}

	// 读取整个响应体。
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 初始化一个变量来存放响应数据结构。
	var response DeviceConfigListResponse

	// 解析响应中的JSON数据，并填充到我们的结构体中。
	if err := json.Unmarshal(p, &response); err != nil {
		return nil, fmt.Errorf("解析响应JSON失败: %w; 响应内容: %s", err, string(p))
	}

	// 返回填充后的响应结构体。
	return &response, nil
}
