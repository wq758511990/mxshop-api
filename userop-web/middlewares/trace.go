package middlewares

import "github.com/gin-gonic/gin"

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
