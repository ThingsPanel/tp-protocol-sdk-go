package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 请求结构体
type DeviceConfigRequest struct {
	DeviceID    string `json:"DeviceId"`
	AccessToken string `json:"AccessToken"`
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
	AccessToken  string      `json:"AccessToken"`
	DeviceType   string      `json:"DeviceType"`
	ID           string      `json:"Id"`
	ProtocolType string      `json:"ProtocolType"`
	SubDevices   []SubDevice `json:"SubDevices"`
}

// 响应结构体
type DeviceConfigResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    DeviceConfigResponseData `json:"data"`
}

// 获取设备配置（通过设备id或者设备令牌）
func (a *API) GetDeviceConfig(request DeviceConfigRequest) (*DeviceConfigResponse, error) {
	resp, err := a.doPostRequest(fmt.Sprintf("%s/api/plugin/device/config", a.BaseURL), request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get device config: status code %d", resp.StatusCode)
	}

	var response DeviceConfigResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
