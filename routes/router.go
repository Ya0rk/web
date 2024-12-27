package routes

import (
	"github.com/gin-gonic/gin"
	v1 "web/api/v1"
	"web/utils/config"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// 管理路由组
	router := r.Group("api/v1")
	{
		router.POST("/register", v1.RegisterApi)
		router.POST("/login", v1.LoginApi)
	}
	r.Run(config.HttpPort)
}
