package controllers

import (
	"fmt"
	"gin-server-template/entity"
	"gin-server-template/models"
	"gin-server-template/services/setting"
	"gin-server-template/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	baseController
}

func (auth *AuthController) Login(ctx *gin.Context) {

	//name := ctx.DefaultPostForm("name", "李小花")
	login := entity.LoginStruct{}
	if err := ctx.ShouldBindBodyWith(&login, binding.JSON); err != nil {
		logs.Error("ShouldBindBodyWith", err.Error())
		auth.ReturnFail(ctx, err.Error())
		return
	}

	logs.Debug("admin password", login)
	// 验证超级密码
	if login.Username == "admin" && login.Password == setting.AdminPassword {
		// 超级admin

		newToken := utils.RandomTimeString(6)
		if err := models.ModToken.AddOne(newToken, 0, setting.TokenKeepTime); err != nil {
			logs.Error("ModToken.AddOne", err.Error())
			auth.ReturnFail(ctx, err.Error())
			return
		}

		tokenInfo, err := models.ModToken.GetOne(newToken)
		if err != nil {
			logs.Error("ModToken.GetOne", err.Error())
			auth.ReturnFail(ctx, err.Error())
			return
		}

		auth.ReturnSucc(ctx, tokenInfo)
	} else {
		///////////////////////////
		// 一般用户
		//用户存在吗
		user_info, err := models.ModUser.GetOneByUsername(login.Username)
		if err != nil {
			logs.Error("ModUser.GetOneByUsername", err.Error())
			auth.ReturnFail(ctx, "username or password is error: "+err.Error())
			return
		}

		// 验证密码
		if user_info.Password != login.Password {
			logs.Error("user_info.Password != login.Password :", user_info)
			auth.ReturnFail(ctx, "username or password is error !")
			return
		}

		newToken := utils.RandomTimeString(6)

		fmt.Println("newToken ==>", newToken, user_info.Id, user_info.Username)

		if err := models.ModToken.AddOne(newToken, user_info.Id, setting.TokenKeepTime); err != nil {
			logs.Error("ModToken.AddOne", err.Error())
			auth.ReturnFail(ctx, err.Error())
			return
		}

		tokenInfo, err := models.ModToken.GetOne(newToken)
		if err != nil {
			logs.Error("ModToken.GetOne", err.Error())
			auth.ReturnFail(ctx, err.Error())
			return
		}

		auth.ReturnSucc(ctx, tokenInfo)
	}
}

func (auth *AuthController) Logout(ctx *gin.Context) {

	xtoken := ctx.Request.Header.Get("x-token")

	if err := models.ModToken.DeleteOne(xtoken); err != nil {
		logs.Warn("ModToken.DeleteOne error:", err.Error())
	} else {
		logs.Debug("logout:", xtoken)
	}
	auth.ReturnSucc(ctx, "ok")
}

func (auth *AuthController) Info(ctx *gin.Context) {
	userinfo, _ := ctx.Get("userinfo")
	auth.ReturnSucc(ctx, userinfo)
}
