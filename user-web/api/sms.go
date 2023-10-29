package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"mxshop-api/user-web/global"
	"net/http"
	"time"
)

func SendSms(ctx *gin.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), "18156566666", "555555", 60*5*time.Second)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "短信发送成功",
	})
}
