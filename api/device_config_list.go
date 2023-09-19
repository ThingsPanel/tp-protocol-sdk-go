package api

import (
	"encoding/json"
	"fmt"
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

func (a *API) GetDeviceConfigList(request DeviceConfigListRequest) (*DeviceConfigListResponse, error) {
	resp, err := a.doPostRequest(fmt.Sprintf("%s/api/plugin/all_device/config", a.BaseURL), request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get device config: status code %d", resp.StatusCode)
	}

	var response DeviceConfigListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
