package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,  //绑定路由规则
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()
}
