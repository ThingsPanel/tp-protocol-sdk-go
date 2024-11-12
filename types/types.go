package types

// Device 设备信息结构体
type Device struct {
	ID           string                 `json:"id"`
	Voucher      string                 `json:"voucher"`
	DeviceNumber string                 `json:"device_number"`
	DeviceType   string                 `json:"device_type"`
	ProtocolType string                 `json:"protocol_type"`
	SubDevices   []SubDevice            `json:"sub_devices,omitempty"`
	Config       map[string]interface{} `json:"config"`
}

// SubDevice 子设备信息结构体
type SubDevice struct {
	DeviceID      string                 `json:"device_id"`
	Voucher       string                 `json:"voucher"`
	DeviceNumber  string                 `json:"device_number"`
	SubDeviceAddr string                 `json:"sub_device_addr"`
	Config        map[string]interface{} `json:"config"`
}

// ServiceAccess 服务接入信息结构体
type ServiceAccess struct {
	ServiceAccessID   string   `json:"service_access_id"`
	ServiceIdentifier string   `json:"service_identifier"`
	Voucher           string   `json:"voucher"`
	Description       string   `json:"description"`
	Remark            string   `json:"remark"`
	Devices           []Device `json:"devices"`
}

// CommonResponse 通用响应结构体
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ServiceAccess 服务接入点列表返回的服务接入点信息
type ServiceAccessRsp struct {
	Description string      `json:"description"`
	Devices     []DeviceRsp `json:"devices"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Remark      string      `json:"remark"`
	Voucher     string      `json:"voucher"`
}

// Device 设备信息
type DeviceRsp struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Voucher         string `json:"voucher"`
	IsEnabled       string `json:"is_enabled"`
	DeviceNumber    string `json:"device_number"`
	SubDeviceAddr   string `json:"sub_device_addr"`
	AdditionalInfo  string `json:"additional_info"`
	ProtocolConfig  string `json:"protocol_config"`
	DeviceConfigID  string `json:"device_config_id"`
	IsOnline        int    `json:"is_online"`
	ServiceAccessID string `json:"service_access_id"`
}
