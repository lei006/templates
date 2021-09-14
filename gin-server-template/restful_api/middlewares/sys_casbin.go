package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
