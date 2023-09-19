package main

import (
	"fmt"
	"net/http"

	tpprotocolsdkgo "github.com/ThingsPanel/tp-protocol-sdk-go"
)

func main() {
	handler := &tpprotocolsdkgo.Handler{
		OnCreateDevice: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handle create notification\n")
		},
		OnUpdateDevice: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handle update notification\n")
		},
		OnDeleteDevice: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handle delete notification\n")
		},
		OnGetForm: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handle get form\n")
		},
	}

	if err := handler.ListenAndServe(":9999"); err != nil {
		panic(err)
	}
}
