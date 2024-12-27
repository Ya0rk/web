package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

func LoginApi(c *gin.Context) {
	var data model.User
	var code int

	err := c.ShouldBindJSON(&data)
	if err != nil {
		panic("Binding Error")
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

	code = service.Login(data.Username, data.Password)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
