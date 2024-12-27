package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

var code int

func SendVerCodeApi(c *gin.Context) {
	var data model.UserEmail

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic("Binding Error")
	}

	if ok := service.IsValidEmail(data.Email); !ok {
		code = errmsg.ERROR_EMAIL_TYPE
		goto Response
	}

	// 发送验证码，并将其存储在cache中，方便验证使用
	code = service.NewEmailService().SendVerificationCode(data.Email)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
