package demo

import (
	"io"
	"log"
	"net/http"

	"../../../internal/biz"
)

func DemoInit() {
	mux := http.NewServeMux()
	mux.HandleFunc("/service/api/v1/getdevice", http.HandlerFunc(httpHandle))

	err := http.ListenAndServe(":18081", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func httpHandle(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		device, _ := biz.GetDevice()
		io.WriteString(res, string(device))
	} else {
		io.WriteString(res, req.RemoteAddr)
	}
}
