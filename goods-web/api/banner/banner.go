package banner

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"mxshop-api/goods-web/api"
	"mxshop-api/goods-web/forms"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"mxshop-api/goods-web/utils"
	"net/http"
)

func List(ctx *gin.Context) {
	rsp, err := global.GoodsSrvClient.BannerList(context.Background(), &emptypb.Empty{})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp.Data))
}

func New(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	rsp, err := global.GoodsSrvClient.CreateBanner(context.Background(), &proto.BannerRequest{
		Image: bannerForm.Image,
		Index: bannerForm.Index,
		Url:   bannerForm.Url,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK.WithData(rsp.Id))
}
func Update(ctx *gin.Context) {
	bannerForm := forms.BannerForm{}
	if err := ctx.ShouldBindJSON(&bannerForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.UpdateBanner(context.Background(), &proto.BannerRequest{
		Id:    bannerForm.ID,
		Image: bannerForm.Image,
		Index: bannerForm.Index,
		Url:   bannerForm.Url,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.OK)
}
func Delete(ctx *gin.Context) {
	delForm := map[string]int32{}
	if err := ctx.ShouldBindJSON(&delForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}
	_, err := global.GoodsSrvClient.DeleteBanner(context.Background(), &proto.BannerRequest{Id: delForm["id"]})
	if err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.OK)
}
