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
		GoodsRouter.GET("stocks", goods.Stocks)
		GoodsRouter.POST("create", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
		GoodsRouter.POST("update", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.POST("updateStatus", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
		GoodsRouter.POST("delete", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)
		GoodsRouter.GET("detail/:id", middlewares.JWTAuth(), goods.GoodsDetail)
	}
}
