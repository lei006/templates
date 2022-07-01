package middlewares

import (
	"gin-server-template/services/setting"
	"gin-server-template/utils/jsonret"
	"gin-server-template/utils/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtToken *jwt.JwtToken

func init() {
	jwtToken = jwt.MakeJwtToken(time.Duration(setting.TokenKeepTime)*time.Hour, setting.AppIssuer, setting.AppName)
}

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			jsonret.ReturnFail(ctx, "请求未携带token，无权限访问")
			return
		}

		claims, err := jwtToken.ParseToken(token)
		if err != nil {
			jsonret.ReturnFail(ctx, "token 非法或者已经过期")
			return
		}
		ctx.Set("authinfo", claims)

		//logs.Debug("authinfo =", claims)

		ctx.Next()
	}
}
