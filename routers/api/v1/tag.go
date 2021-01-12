package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"golang.org/x/crypto/ocsp"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	code := e.SUCCESS
	m := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		m["name"] = name
	}

	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(state).MustInt()
		m["state"] = state
	}

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, m)
	data["total"] = models.GetCount(m)

	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}

func AddTag(c *gin.Context) {

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}