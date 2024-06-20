package api

// 子设备信息结构体
type SubDevice struct {
	DeviceID               string                 `json:"device_id"`
	Voucher                string                 `json:"voucher"`
	DeviceNumber           string                 `json:"device_number"`
	SubDeviceAddr          string                 `json:"sub_device_addr"`
	Config                 map[string]interface{} `json:"config"`
	ProtocolConfigTemplate map[string]interface{} `json:"protocol_config_template"` // 子设备配置的protocol_config（表单数据）
}

// 设备信息结构体
type Device struct {
	ID                     string                 `json:"id"`
	Voucher                string                 `json:"voucher"`
	DeviceNumber           string                 `json:"device_number"`
	DeviceType             string                 `json:"device_type"`
	ProtocolType           string                 `json:"Protocol_type"`
	SubDevices             []SubDevice            `json:"sub_devices"`
	Config                 map[string]interface{} `json:"config"`
	ProtocolConfigTemplate map[string]interface{} `json:"protocol_config_template"` // 子设备配置的protocol_config（表单数据）
}

// 服务接入结构体
type ServiceAccess struct {
	ServiceAccessID             string   `json:"service_access_id"`
	ServiceIdentifier           string   `json:"service_identifier"`
	Voucher                     string   `json:"voucher"`                        // 服务凭证（表单数据）
	ServiceAccessConfigTemplate string   `json:"service_access_config_template"` // 服务接入配置
	Description                 string   `json:"description"`
	Remark                      string   `json:"remark"`
	Devices                     []Device `json:"devices"`
}
