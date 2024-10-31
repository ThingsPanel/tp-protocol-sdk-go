// handler/handler.go

package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// HandlerConfig 处理器配置
type HandlerConfig struct {
	Logger *log.Logger // 日志记录器
}

// Handler 回调处理器
type Handler struct {
	logger                  *log.Logger
	formConfigHandler       func(req *GetFormConfigRequest) (interface{}, error)
	deviceDisconnectHandler func(req *DeviceDisconnectRequest) error
	notificationHandler     func(req *NotificationRequest) error
	getDeviceListHandler    func(req *GetDeviceListRequest) (*DeviceListResponse, error)
}

// NewHandler 创建一个新的处理器实例
func NewHandler(config HandlerConfig) *Handler {
	logger := config.Logger
	if logger == nil {
		logger = log.New(log.Writer(), "[TP-Handler] ", log.LstdFlags|log.Lshortfile)
	}

	return &Handler{
		logger: logger,
	}
}

// parseQueryParams 解析查询参数
func parseQueryParams(r *http.Request, obj interface{}) error {
	// 这里可以使用反射来实现，为了简单起见，我们针对具体类型处理
	switch v := obj.(type) {
	case *GetFormConfigRequest:
		v.ProtocolType = r.URL.Query().Get("protocol_type")
		v.DeviceType = r.URL.Query().Get("device_type")
		v.FormType = r.URL.Query().Get("form_type")

		// 验证必填参数
		if v.ProtocolType == "" || v.DeviceType == "" || v.FormType == "" {
			return fmt.Errorf("missing required query parameters")
		}

	case *GetDeviceListRequest:
		v.Voucher = r.URL.Query().Get("voucher")
		v.ServiceIdentifier = r.URL.Query().Get("service_identifier")

		pageSize := r.URL.Query().Get("page_size")
		page := r.URL.Query().Get("page")

		if v.Voucher == "" || v.ServiceIdentifier == "" || pageSize == "" || page == "" {
			return fmt.Errorf("missing required query parameters")
		}

		var err error
		v.PageSize, err = strconv.Atoi(pageSize)
		if err != nil {
			return fmt.Errorf("invalid page_size")
		}

		v.Page, err = strconv.Atoi(page)
		if err != nil {
			return fmt.Errorf("invalid page")
		}
	}
	return nil
}

func (h *Handler) handleFormConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req GetFormConfigRequest
	if err := parseQueryParams(r, &req); err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.formConfigHandler(&req)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, "success", data)
}

func (h *Handler) handleDeviceDisconnect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req DeviceDisconnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.deviceDisconnectHandler(&req); err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, "success", nil)
}

func (h *Handler) handleNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.notificationHandler(&req); err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, "success", nil)
}

func (h *Handler) handleGetDeviceList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req GetDeviceListRequest
	if err := parseQueryParams(r, &req); err != nil {
		h.writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.getDeviceListHandler(&req)
	if err != nil {
		h.writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.writeResponse(w, http.StatusOK, "success", resp.Data)
}

func (h *Handler) writeError(w http.ResponseWriter, code int, message string) {
	h.writeResponse(w, code, message, nil)
}

func (h *Handler) writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp := CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)
}

// ServeHTTP 实现 http.Handler 接口
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("收到请求: %s %s", r.Method, r.URL.Path)

	switch r.URL.Path {
	case "/api/v1/form/config":
		h.handleFormConfig(w, r)
	case "/api/v1/device/disconnect":
		h.handleDeviceDisconnect(w, r)
	case "/api/v1/plugin/notification":
		h.handleNotification(w, r)
	case "/api/v1/plugin/device/list":
		h.handleGetDeviceList(w, r)
	default:
		http.NotFound(w, r)
	}
}

// SetFormConfigHandler 设置表单配置处理函数
func (h *Handler) SetFormConfigHandler(handler func(req *GetFormConfigRequest) (interface{}, error)) {
	h.formConfigHandler = handler
}

// SetDeviceDisconnectHandler 设置设备断开处理函数
func (h *Handler) SetDeviceDisconnectHandler(handler func(req *DeviceDisconnectRequest) error) {
	h.deviceDisconnectHandler = handler
}

// SetNotificationHandler 设置通知处理函数
func (h *Handler) SetNotificationHandler(handler func(req *NotificationRequest) error) {
	h.notificationHandler = handler
}

// SetGetDeviceListHandler 设置获取设备列表处理函数
func (h *Handler) SetGetDeviceListHandler(handler func(req *GetDeviceListRequest) (*DeviceListResponse, error)) {
	h.getDeviceListHandler = handler
}

// Start 启动HTTP服务
func (h *Handler) Start(addr string) error {
	h.logger.Printf("启动HTTP服务: %s", addr)
	return http.ListenAndServe(addr, h)
}
