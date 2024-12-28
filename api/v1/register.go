package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/middleware"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

func RegisterApi(c *gin.Context) {
	var data model.UserRegister
	var code int
	var token string

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic("Binding Error")
	}

	// 一些检查
	code = service.Check(data)
	if code != errmsg.SUCCESS {
		goto Response
	}

	// 检验 验证码 的正确性
	if ok := service.NewEmailService().VerifyVerificationCode(data.Email, data.VerificationCode); !ok {
		code = errmsg.ERROR_VERIFICATIONCODE
	}

Response:
	if code == errmsg.SUCCESS {
		// 注册之后创建新用户
		code = service.CreateUser(&data)
		token, code = middleware.SetToken(data.Username)
	} else {
		token = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"token":   token,
		"message": errmsg.GetErrMsg(code),
	})
}
