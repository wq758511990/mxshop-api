package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/category"
)

func InitCategoryRouter(r *gin.RouterGroup) {
	categoryRouter := r.Group("category")
	{
		categoryRouter.GET("list", category.List)
		categoryRouter.GET("detail", category.Detail)
		categoryRouter.POST("create", category.New)
		categoryRouter.POST("delete", category.Delete)
		categoryRouter.POST("update", category.Update)
	}
}
