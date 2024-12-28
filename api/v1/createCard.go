package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web/model"
	"web/service"
	"web/utils/errmsg"
)

func CreateCardApi(c *gin.Context) {
	var data model.UserCard
	var code int
	if err := c.ShouldBind(&data); err != nil {
		panic(err.Error())
	}

	// 检查
	if err := data.Validate(); err != nil {
		code = errmsg.ERROR_CARD_MSG
		goto Response
	}

	// 加入数据库
	code = service.NewCardService().CreateCard(data)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCardsApi(c *gin.Context) {
	var code int
	var data []model.UserCard
	var num int64
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize <= 0 || pageNum <= 0 {
		code = errmsg.ERROR_PARAM
		data = []model.UserCard{}
		num = 0
		goto Response
	}

	// 返回所有的名片
	data, code, num = service.NewCardService().GetCards(pageSize, pageNum)

Response:
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"total":   num,
		"message": errmsg.GetErrMsg(code),
	})
}
