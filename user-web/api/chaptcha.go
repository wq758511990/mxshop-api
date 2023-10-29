package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误, : %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "验证码生成失败",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64,
	})
}
