package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/userop-web/api/address"
	"mxshop-api/userop-web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)
	}
}
