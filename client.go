package tpprotocolsdkgo

import (
	"github.com/ThingsPanel/tp-protocol-sdk-go/api"
)

// 用于访问ThingsPanel API的客户端
type Client struct {
	API *api.API
}

// 创建一个新的客户端实例
func NewClient(baseURL string) *Client {
	return &Client{
		API: api.NewAPI(baseURL),
	}
}
