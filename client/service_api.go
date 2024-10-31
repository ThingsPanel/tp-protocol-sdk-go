// client/service_api.go

package client

import (
	"context"
	"fmt"

	"github.com/ThingsPanel/tp-protocol-sdk-go/types"
)

// ServiceAPI 服务接入相关API封装
type ServiceAPI struct {
	client *APIClient
}

// ServiceAccessRequest 获取服务接入点请求
type ServiceAccessRequest struct {
	ServiceAccessID   string `json:"service_access_id"`
	ServiceIdentifier string `json:"service_identifier"`
}

// ServiceAccessResponse 服务接入响应
type ServiceAccessResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    types.ServiceAccess `json:"data"`
}

// HeartbeatRequest 心跳请求
type HeartbeatRequest struct {
	ServiceIdentifier string `json:"service_identifier"`
}

// HeartbeatResponse 心跳响应
type HeartbeatResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewServiceAPI 创建服务API实例
func NewServiceAPI(client *APIClient) *ServiceAPI {
	return &ServiceAPI{
		client: client,
	}
}

// GetServiceAccess 获取服务接入点信息
func (s *ServiceAPI) GetServiceAccess(ctx context.Context, req *ServiceAccessRequest) (*ServiceAccessResponse, error) {
	s.client.logger.Printf("开始获取服务接入点信息: serviceAccessID=%s", req.ServiceAccessID)

	var resp ServiceAccessResponse
	err := s.client.Post(ctx, "/api/v1/plugin/service/access", req, &resp)
	if err != nil {
		s.client.logger.Printf("获取服务接入点信息失败: %v", err)
		return nil, fmt.Errorf("获取服务接入点信息失败: %w", err)
	}

	s.client.logger.Printf("获取服务接入点信息成功: serviceIdentifier=%s",
		resp.Data.ServiceIdentifier)
	return &resp, nil
}

// SendHeartbeat 发送服务心跳
func (s *ServiceAPI) SendHeartbeat(ctx context.Context, req *HeartbeatRequest) (*HeartbeatResponse, error) {
	s.client.logger.Printf("开始发送服务心跳: serviceIdentifier=%s", req.ServiceIdentifier)

	var resp HeartbeatResponse
	err := s.client.Post(ctx, "/api/v1/plugin/heartbeat", req, &resp)
	if err != nil {
		s.client.logger.Printf("发送服务心跳失败: %v", err)
		return nil, fmt.Errorf("发送服务心跳失败: %w", err)
	}

	s.client.logger.Printf("发送服务心跳成功")
	return &resp, nil
}
