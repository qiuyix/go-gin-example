package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/setting"
	"net/http"
)

func main() {
	router := gin.Default()
	
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,  //绑定路由规则
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe()
}
