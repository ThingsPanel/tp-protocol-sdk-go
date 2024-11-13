// handler/types.go

package handler

// GetFormConfigRequest 获取表单配置请求
type GetFormConfigRequest struct {
	ProtocolType string `form:"protocol_type" binding:"required"` // 协议/服务标识符
	DeviceType   string `form:"device_type,omitempty"`            // 1-设备 2-网关 3-子设备
	FormType     string `form:"form_type" binding:"required"`     // CFG-配置表单 VCR-凭证表单 SVCR-服务凭证表单
}

// DeviceDisconnectRequest 设备断开连接请求
type DeviceDisconnectRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
}

// NotificationRequest 通知请求
type NotificationRequest struct {
	MessageType string `json:"message_type" binding:"required"` // 1-服务配置修改
	Message     string `json:"message,omitempty"`               // 消息内容
}

// GetDeviceListRequest 获取设备列表请求
type GetDeviceListRequest struct {
	Voucher           string `form:"voucher" binding:"required"`            // 凭证
	ServiceIdentifier string `form:"service_identifier" binding:"required"` // 服务标识符
	PageSize          int    `form:"page_size" binding:"required"`          // 每页数量
	Page              int    `form:"page" binding:"required"`               // 页数
}

// DeviceItem 设备列表项
type DeviceItem struct {
	DeviceName   string `json:"device_name"`
	Description  string `json:"description"`
	DeviceNumber string `json:"device_number"`
}

// DeviceListData 设备列表数据
type DeviceListData struct {
	List  []DeviceItem `json:"list"`
	Total int          `json:"total"`
}

// CommonResponse 通用响应结构
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// DeviceListResponse 设备列表响应
type DeviceListResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    DeviceListData `json:"data"`
}
