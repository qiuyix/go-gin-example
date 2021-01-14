package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	code := e.SUCCESS
	m := make(map[string]interface{})    // 查询map
	data := make(map[string]interface{}) // 返回map

	if name != "" {
		m["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		m["state"] = state
	}

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, m)
	data["total"] = models.GetCount(m)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("create_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createBy, "created_by").Message("创建人蹦年为空")
	valid.MaxSize(createBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
