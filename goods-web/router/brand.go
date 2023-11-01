package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/brand"
)

func InitBrandRouter(r *gin.RouterGroup) {
	brandRouter := r.Group("brand")
	{
		brandRouter.GET("list", brand.List)
		brandRouter.POST("create", brand.New)
		brandRouter.POST("delete", brand.Delete)
		brandRouter.POST("update", brand.Update)
	}
	categoryBrandRouter := r.Group("categoryBrand")
	{
		categoryBrandRouter.POST("create", brand.NewCategoryBrandList)
		categoryBrandRouter.GET("detail", brand.GetCategoryBrandList)
		categoryBrandRouter.GET("list", brand.CategoryBrandList)
		categoryBrandRouter.POST("update", brand.UpdateCategoryBrand)
		categoryBrandRouter.POST("delete", brand.DeleteCategoryBrand)
	}
}
