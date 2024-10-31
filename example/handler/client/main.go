// examples/handler/client/main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[Client] ", log.LstdFlags|log.Lshortfile)
	baseURL := "http://localhost:8080"

	// 1. 测试获取表单配置
	logger.Println("测试获取表单配置...")
	formConfig := testGetFormConfig(baseURL, logger)
	logger.Printf("表单配置: %+v\n", formConfig)

	// 2. 测试设备断开连接
	logger.Println("\n测试设备断开连接...")
	testDeviceDisconnect(baseURL, "device-001", logger)

	// 3. 测试发送通知
	logger.Println("\n测试发送通知...")
	testSendNotification(baseURL, logger)

	// 4. 测试获取设备列表
	logger.Println("\n测试获取设备列表...")
	deviceList := testGetDeviceList(baseURL, logger)
	logger.Printf("设备列表: %+v\n", deviceList)
}

// 获取表单配置
func testGetFormConfig(baseURL string, logger *log.Logger) map[string]interface{} {
	// 构建请求URL和参数
	params := url.Values{}
	params.Add("protocol_type", "modbus")
	params.Add("device_type", "1")
	params.Add("form_type", "CFG")

	url := fmt.Sprintf("%s/api/v1/form/config?%s", baseURL, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Fatalf("解析响应失败: %v", err)
	}

	return result
}

// 测试设备断开连接
func testDeviceDisconnect(baseURL string, deviceID string, logger *log.Logger) {
	reqBody, err := json.Marshal(map[string]string{
		"device_id": deviceID,
	})
	if err != nil {
		logger.Fatalf("构建请求体失败: %v", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/device/disconnect", baseURL),
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		logger.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalf("读取响应失败: %v", err)
	}

	logger.Printf("响应: %s", string(body))
}

// 测试发送通知
func testSendNotification(baseURL string, logger *log.Logger) {
	notifyData := map[string]string{
		"message_type": "1",
		"message":      `{"service_access_id":"service-001"}`,
	}

	reqBody, err := json.Marshal(notifyData)
	if err != nil {
		logger.Fatalf("构建请求体失败: %v", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/plugin/notification", baseURL),
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		logger.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalf("读取响应失败: %v", err)
	}

	logger.Printf("响应: %s", string(body))
}

// 测试获取设备列表
func testGetDeviceList(baseURL string, logger *log.Logger) map[string]interface{} {
	// 构建请求URL和参数
	params := url.Values{}
	params.Add("voucher", "test-voucher")
	params.Add("service_identifier", "test-service")
	params.Add("page_size", "10")
	params.Add("page", "1")

	url := fmt.Sprintf("%s/api/v1/plugin/device/list?%s", baseURL, params.Encode())

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Fatalf("解析响应失败: %v", err)
	}

	return result
}
