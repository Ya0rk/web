package routes

import (
	"github.com/gin-gonic/gin"
	v1 "web/api/v1"
	"web/middleware"
	"web/utils/config"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())

	// 管理路由组
	// 需要鉴权
	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{

	}

	// 不需要鉴权
	public := r.Group("api/v1")
	{
		public.POST("register", v1.RegisterApi)
		public.POST("login/passwd", v1.LoginByPasswdApi)
		public.POST("login/sendVerCode", v1.SendVerCodeApi)
		public.POST("login/email", v1.LoginByEmailApi)
		public.POST("recover/passwd", v1.RecoverPasswdApi)
	}
	r.Run(config.HttpPort)
}
