package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 请求结构体
type DeviceConfigRequest struct {
	DeviceID     string `json:"DeviceId"`
	AccessToken  string `json:"AccessToken"`
	DeviceNumber string `json:"DeviceNumber"`
}

// 子设备信息结构体
type SubDevice struct {
	AccessToken   string                 `json:"AccessToken"`
	DeviceID      string                 `json:"DeviceId"`
	SubDeviceAddr string                 `json:"SubDeviceAddr"`
	Config        map[string]interface{} `json:"Config"`
}

// 设备信息结构体
type DeviceConfigResponseData struct {
	AccessToken  string                 `json:"AccessToken"`
	DeviceType   string                 `json:"DeviceType"`
	ID           string                 `json:"Id"`
	ProtocolType string                 `json:"ProtocolType"`
	SubDevices   []SubDevice            `json:"SubDevices"`
	DeviceConfig map[string]interface{} `json:"DeviceConfig"`
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
	apiEndpoint := fmt.Sprintf("%s/api/plugin/device/config", a.BaseURL)

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
	p, err := ioutil.ReadAll(resp.Body)
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
