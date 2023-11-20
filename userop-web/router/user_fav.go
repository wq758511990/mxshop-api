package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/userop-web/api/user_fav"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	UserFavRouter := Router.Group("userfav")
	{
		UserFavRouter.GET("", user_fav.List)
	}
}
