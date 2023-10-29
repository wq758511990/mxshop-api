package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/goods"
	"mxshop-api/user-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")

	{
		GoodsRouter.GET("list", goods.List)
		GoodsRouter.POST("create", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
	}
}
