package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/userop-web/api/message"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	messageGroup := Router.Group("message")
	{
		messageGroup.GET("", message.List)
	}
}
