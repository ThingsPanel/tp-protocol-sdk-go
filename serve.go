package tpprotocolsdkgo

import (
	"net/http"
)

// Handler 结构体用于存储用户提供的回调函数
type Handler struct {
	// 新增设备回调
	OnCreateDevice func(w http.ResponseWriter, r *http.Request)
	// 更新设备回调
	OnUpdateDevice func(w http.ResponseWriter, r *http.Request)
	// 删除设备回调
	OnDeleteDevice func(w http.ResponseWriter, r *http.Request)
	// 获取协议插件的json表单
	OnGetForm func(w http.ResponseWriter, r *http.Request)
}

// ListenAndServe 函数启动一个HTTP服务器来处理TP平台的通知
func (h *Handler) ListenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/device/config/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.OnCreateDevice(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/device/config/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.OnUpdateDevice(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/device/config/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.OnDeleteDevice(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/form/config", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.OnGetForm(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	return http.ListenAndServe(addr, mux)
}
