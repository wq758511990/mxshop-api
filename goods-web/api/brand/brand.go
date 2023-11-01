package brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api"
	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"mxshop-api/goods-web/utils"
	"net/http"
	"strconv"
)

func New(ctx *gin.Context) {
	brandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&brandForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(map[string]interface{}{
		"id": rsp.Id,
	}))
}

func Delete(ctx *gin.Context) {
	delForm := map[string]int32{}
	if err := ctx.ShouldBindJSON(&delForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.DeleteBrand(context.Background(), &proto.BrandRequest{Id: delForm["id"]})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}

func Update(ctx *gin.Context) {
	updateBrandForm := forms.BrandForm{}
	if err := ctx.ShouldBindJSON(&updateBrandForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   updateBrandForm.ID,
		Name: updateBrandForm.Name,
		Logo: updateBrandForm.Logo,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp))
}

func List(ctx *gin.Context) {
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pnInt),
		PagePerNums: int32(pSizeInt),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp))
}

func GetCategoryBrandList(ctx *gin.Context) {
	id := ctx.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BadRequest)
		return
	}
	rsp, err := global.GoodsSrvClient.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: int32(i),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp.Data))
}

func NewCategoryBrandList(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp.Id))
}

func CategoryBrandList(ctx *gin.Context) {
	rsp, err := global.GoodsSrvClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, item := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = item.Id
		reMap["category"] = map[string]interface{}{
			"id":   item.Category.Id,
			"name": item.Category.Name,
		}
		reMap["brand"] = map[string]interface{}{
			"id":   item.Brand.Id,
			"name": item.Brand.Name,
			"logo": item.Brand.Logo,
		}
		result = append(result, reMap)
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(result))
}

func UpdateCategoryBrand(ctx *gin.Context) {
	categoryBrandForm := forms.CategoryBrandForm{}
	if err := ctx.ShouldBindJSON(&categoryBrandForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         categoryBrandForm.ID,
		CategoryId: int32(categoryBrandForm.CategoryId),
		BrandId:    int32(categoryBrandForm.BrandId),
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}

func DeleteCategoryBrand(ctx *gin.Context) {
	delForm := map[string]int32{}
	if err := ctx.ShouldBindJSON(&delForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id: delForm["id"],
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}
