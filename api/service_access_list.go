package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 请求结构体
type ServiceAccessListRequest struct {
	ServiceIdentifier string `json:"service_identifier"`
}

// 响应结构体
type ServiceAccessListResponseData struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []ServiceAccess `json:"data"`
}

// 获取服务接入点列表
func (a *API) GetServiceAccessList(request ServiceAccessListRequest) (*ServiceAccessListResponseData, error) {

	// 构建API终端的URL
	apiEndpoint := fmt.Sprintf("%s/api/v1/plugin/service/access/list", a.BaseURL)

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
		return nil, fmt.Errorf("获取服务接入点列表失败: HTTP状态码 %d", resp.StatusCode)
	}

	// 读取整个响应体。
	p, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 解析json响应
	var data ServiceAccessListResponseData
	err = json.Unmarshal(p, &data)
	if err != nil {
		return nil, fmt.Errorf("解析json响应失败: %w", err)
	}

	return &data, nil
}
