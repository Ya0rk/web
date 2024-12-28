package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

// 使用邮箱找回密码
// 1. 验证邮箱是否存在
// 2. 设置新密码
// 3. 发送验证码
// 4. 验证
// 5，修改
func RecoverPasswdApi(c *gin.Context) {
	var data model.RecoverPasswd
	var code int

	err := c.ShouldBind(&data)
	if err != nil {
		panic(err.Error())
	}

	// 一些检查
	switch {
	case !service.IsValidEmail(data.Email):
		code = errmsg.ERROR_EMAIL_TYPE // 如果邮箱无效，返回邮箱错误
		goto Response
	case len(data.Password) < 8 || len(data.Password) > 24:
		code = errmsg.ERROR_PASSWORD_LEN // 如果密码长度不符合要求，返回密码长度错误
		goto Response
	case len(data.VerificationCode) != 6:
		code = errmsg.ERROR_VERIFICATIONCODE_LEN // 如果验证码长度小于6，返回验证码长度错误
		goto Response
	}

	code, _ = service.CheckEmailCode(data.Email, data.VerificationCode)
	if code == errmsg.SUCCESS {
		code = service.RecoverPasswd(data.Email, data.Password)
	}

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
