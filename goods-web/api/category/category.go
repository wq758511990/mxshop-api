package category

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"mxshop-api/goods-web/api"
	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"mxshop-api/goods-web/utils"
	"net/http"
	"strconv"
)

func List(ctx *gin.Context) {
	rsp, err := global.GoodsSrvClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	data := make([]interface{}, 0)
	err = json.Unmarshal([]byte(rsp.JsonData), &data)
	if err != nil {
		zap.S().Errorw("[List] 查询 【分类列表】失败： ", err.Error())
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(data))
}

func Detail(ctx *gin.Context) {
	id := ctx.Query("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BadRequest)
		return
	}

	reMap := map[string]interface{}{}
	subCategorys := make([]interface{}, 0)
	if r, err := global.GoodsSrvClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: int32(i),
	}); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	} else {
		for _, category := range r.SubCategorys {
			subCategorys = append(subCategorys, map[string]interface{}{
				"id":              category.Id,
				"name":            category.Name,
				"level":           category.Level,
				"parent_category": category.ParentCategory,
				"is_tab":          category.IsTab,
			})
		}
		reMap["id"] = r.Info.Id
		reMap["name"] = r.Info.Name
		reMap["level"] = r.Info.Level
		reMap["parent_category"] = r.Info.ParentCategory
		reMap["is_tab"] = r.Info.IsTab
		reMap["sub_categorys"] = subCategorys

		ctx.JSON(http.StatusOK, utils.OK.WithData(reMap))
	}
	return

}

func New(ctx *gin.Context) {
	categoryForm := &forms.CategoryForm{}

	if err := ctx.ShouldBindJSON(&categoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateCategory(context.Background(), &proto.CategoryInfoRequest{
		ParentCategory: categoryForm.ParentCategory,
		IsTab:          *categoryForm.IsTab,
		Level:          categoryForm.Level,
		Name:           categoryForm.Name,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp))
}

func Delete(ctx *gin.Context) {
	delForm := map[string]int32{}
	if err := ctx.ShouldBindJSON(&delForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.DeleteCategory(context.Background(), &proto.DeleteCategoryRequest{
		Id: delForm["id"],
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}

func Update(ctx *gin.Context) {
	updateCategoryForm := proto.CategoryInfoRequest{}
	if err := ctx.ShouldBindJSON(&updateCategoryForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.UpdateCategory(context.Background(), &updateCategoryForm)
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}
