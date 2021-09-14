package middlewares

import (
	"gin-server-template/models"
	"gin-server-template/services/setting"
	"gin-server-template/utils/jsonret"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			jsonret.ReturnFail(ctx, "请求未携带token，无权限访问")
			return
		}

		tokenInfo, err := models.ModToken.IsExpired(token)
		if err != nil {
			jsonret.ReturnFail(ctx, "Token: 非法或者已经过期")
			return
		}

		// 给token 续期...
		models.ModToken.RenewalOne(token, setting.TokenKeepTime)

		if tokenInfo.UserId == 0 {
			//管理员
			userInfo := models.User{
				Id:       0,
				Username: "admin",
				Nickname: "超级管理员",
				Avatar:   "admin",
			}
			ctx.Set("userinfo", userInfo)
		} else {
			//一般用户
			userInfo, err := models.ModUser.GetOne(tokenInfo.UserId)
			if err != nil {
				logs.Error("TokenAuth: ModUser.GetOne", err.Error())
				jsonret.ReturnFail(ctx, "Token: 非法或者已经过期")
				return
			}
			userInfo.Password = ""

			if userInfo.IsEnable == false {
				logs.Error("TokenAuth: ModUser.GetOne")
				jsonret.ReturnFailMessage(ctx, "该用户已经停用!")
				return
			}

			ctx.Set("userinfo", userInfo)
		}

		ctx.Next()
	}
}
