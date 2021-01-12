package main

import (
	"net/http"
	"time"
)


// 未使用gin的开启http服务
func main1() {
	mux := http.NewServeMux()

	mux.HandleFunc("/text", func(writer http.ResponseWriter, request *http.Request) {

	})

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	server.ListenAndServe()
}