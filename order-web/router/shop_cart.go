package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/shop_cart"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("shop-cart")
	{
		GoodsRouter.GET("list", shop_cart.List)
	}
}
