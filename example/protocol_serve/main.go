package main

import (
	"fmt"
	"net/http"

	tpprotocolsdkgo "github.com/ThingsPanel/tp-protocol-sdk-go"
)

func main() {
	handler := &tpprotocolsdkgo.Handler{
		OnGetForm: func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Handle get form\n")
		},
	}

	if err := handler.ListenAndServe(":9999"); err != nil {
		panic(err)
	}
}
