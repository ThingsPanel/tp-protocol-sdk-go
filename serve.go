package tpprotocolsdkgo

import (
	"net/http"
)

// Handler 结构体用于存储用户提供的回调函数
type Handler struct {
	// 获取协议插件的json表单
	OnGetForm   func(w http.ResponseWriter, r *http.Request)
	OnGetFormV2 func(w http.ResponseWriter, r *http.Request)
	// 断开设备连接回调（让设备重新连接）
	OnDisconnectDevice func(w http.ResponseWriter, r *http.Request)
	// 获取设备列表
	GetDeviceList func(w http.ResponseWriter, r *http.Request)
	// 通知事件
	OnNotifyEvent func(w http.ResponseWriter, r *http.Request)
}

// ListenAndServe 函数启动一个HTTP服务器来处理TP平台的通知
func (h *Handler) ListenAndServe(addr string) error {
	mux := http.NewServeMux()

	// 获取协议插件的json表单
	mux.HandleFunc("/api/v1/form/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.OnGetForm(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	// 获取协议插件的json表单
	mux.HandleFunc("/api/v2/form/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.OnGetFormV2(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// 断开设备连接
	mux.HandleFunc("/api/v1/device/disconnect", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.OnDisconnectDevice(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// 获取设备列表
	mux.HandleFunc("/api/v1/plugin/device/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetDeviceList(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// 通知事件
	mux.HandleFunc("/api/v1/notify/event", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.OnNotifyEvent(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	return http.ListenAndServe(addr, mux)
}
