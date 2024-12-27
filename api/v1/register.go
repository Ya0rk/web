package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

func RegisterApi(c *gin.Context) {
	var data model.UserRegister
	var code int

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic("Binding Error")
	}

	// 一些检查
	code = service.Check(data)
	if code != errmsg.SUCCESS {
		goto Response
	}

	// 向邮箱发送验证信息
	code = service.SendVerifiMsg(data.Email)

Response:
	if code == errmsg.SUCCESS {
		code = service.CreateUser(&data)
	} else {
		data = model.UserRegister{} // 返回一个空的data，避免泄露其他user信息
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
