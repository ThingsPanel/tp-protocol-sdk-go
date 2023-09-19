package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// API 用于访问ThingsPanel API
type API struct {
	BaseURL    string
	httpClient *http.Client
}

// NewAPI 创建一个新的API实例
func NewAPI(baseURL string) *API {
	return &API{
		BaseURL:    baseURL,
		httpClient: &http.Client{}, // 初始化 httpClient
	}
}

// doRequest 发送HTTP-post请求
func (a *API) doPostRequest(url string, reqBody interface{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := a.httpClient.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}

	return resp, nil
}
