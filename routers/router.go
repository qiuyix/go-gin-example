package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/setting"
	v1 "go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New() // git.Default() 也使用git.New()创建engine实例i，但是会默认使用Logger和Recovery中间件，git.New() 则一切自定义

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "ok",
		})
	})

	return r
}
