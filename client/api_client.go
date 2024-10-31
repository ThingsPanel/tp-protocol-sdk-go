// client/api_client.go

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// APIClient TP平台API客户端
type APIClient struct {
	baseURL    string       // API基础URL
	httpClient *http.Client // HTTP客户端
	logger     *log.Logger  // 日志记录器
}

// APIClientOption 定义客户端配置选项
type APIClientOption func(*APIClient)

// WithTimeout 设置超时时间选项
func WithTimeout(timeout time.Duration) APIClientOption {
	return func(c *APIClient) {
		c.httpClient.Timeout = timeout
	}
}

// WithLogger 设置日志记录器选项
func WithLogger(logger *log.Logger) APIClientOption {
	return func(c *APIClient) {
		c.logger = logger
	}
}

// NewAPIClient 创建新的API客户端实例
func NewAPIClient(baseURL string, opts ...APIClientOption) *APIClient {
	client := &APIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: log.New(log.Writer(), "[TP-SDK] ", log.LstdFlags|log.Lshortfile),
	}

	// 应用配置选项
	for _, opt := range opts {
		opt(client)
	}

	client.logger.Printf("初始化API客户端: baseURL=%s", baseURL)
	return client
}

// doRequest 执行HTTP请求并处理响应
func (c *APIClient) doRequest(ctx context.Context, method, path string, reqBody, respBody interface{}) error {
	// 构建完整URL
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	c.logger.Printf("准备发送请求: method=%s, url=%s", method, url)

	// 序列化请求体
	var bodyReader io.Reader
	if reqBody != nil {
		bodyBytes, err := json.Marshal(reqBody)
		if err != nil {
			c.logger.Printf("请求体序列化失败: %v", err)
			return fmt.Errorf("序列化请求体失败: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
		c.logger.Printf("请求体: %s", string(bodyBytes))
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		c.logger.Printf("创建请求失败: %v", err)
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	// 执行请求
	startTime := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Printf("请求执行失败: %v", err)
		return fmt.Errorf("执行请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 记录请求耗时
	c.logger.Printf("请求完成: 耗时=%v, 状态码=%d", time.Since(startTime), resp.StatusCode)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Printf("读取响应体失败: %v", err)
		return fmt.Errorf("读取响应体失败: %w", err)
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		c.logger.Printf("请求返回非200状态码: status=%d, body=%s", resp.StatusCode, string(body))
		return fmt.Errorf("请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	// 解析响应体
	if respBody != nil {
		if err := json.Unmarshal(body, respBody); err != nil {
			c.logger.Printf("响应体解析失败: %v, body=%s", err, string(body))
			return fmt.Errorf("解析响应体失败: %w", err)
		}
		c.logger.Printf("响应体解析成功")
	}

	return nil
}

// Get 执行GET请求
func (c *APIClient) Get(ctx context.Context, path string, response interface{}) error {
	return c.doRequest(ctx, http.MethodGet, path, nil, response)
}

// Post 执行POST请求
func (c *APIClient) Post(ctx context.Context, path string, request, response interface{}) error {
	return c.doRequest(ctx, http.MethodPost, path, request, response)
}
