package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PasswordLogin)
		UserRouter.POST("register", api.Register)
	}
}
