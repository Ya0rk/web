package routes

import (
	"github.com/gin-gonic/gin"
	"web/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// 管理路由组
	r := router.Group("api/v1")
	{
		r.POST("/register")
	}
}
