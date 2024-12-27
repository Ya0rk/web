package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

// 使用密码验证登录
func LoginByPasswdApi(c *gin.Context) {
	var data model.User
	var code int

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic(err.Error())
	}

	// 长度检查
	switch {
	case len(data.Username) < 4 || len(data.Username) > 20:
		code = errmsg.ERROR_USERNAME_LEN // 用户名长度必须在4到20之间
		goto Response
	case len(data.Password) < 8 || len(data.Password) > 24:
		code = errmsg.ERROR_PASSWORD_LEN
		goto Response
	}

	code = service.CheckPasswd(data.Username, data.Password)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 使用邮箱方式登录
func LoginByEmailApi(c *gin.Context) {
	var data model.UserEmailLogin
	var code int

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic("Binding Error")
	}

	// 长度检查
	if ok := service.IsValidEmail(data.Email); !ok {
		code = errmsg.ERROR_EMAIL_TYPE
		goto Response
	}
	if len(data.VerificationCode) != 6 {
		code = errmsg.ERROR_VERIFICATIONCODE_LEN
		goto Response
	}

	code = service.CheckEmailCode(data.Email, data.VerificationCode)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
