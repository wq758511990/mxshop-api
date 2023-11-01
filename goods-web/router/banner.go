package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/banner"
)

func InitBannerRouter(r *gin.RouterGroup) {
	bannerRouter := r.Group("banner")
	{
		bannerRouter.GET("list", banner.List)
		bannerRouter.POST("create", banner.New)
		bannerRouter.POST("delete", banner.Delete)
		bannerRouter.POST("update", banner.Update)
	}
}
