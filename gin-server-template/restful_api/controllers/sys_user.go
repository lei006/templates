package controllers

import (
	"gin-server-template/models"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
	baseController
	NoCheckToken bool //	不检查token
	Prefix       string
}

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
}

func (user *UserController) GetList(ctx *gin.Context) {

	//users := []User{{ID: 123, Name: "张三1", NickName: "11"}, {ID: 456, Name: "李四"}}
	user_list, err := models.ModUser.GetAll()
	if err != nil {
		logs.Error("ModUser.GetAll", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	user_total, err := models.ModUser.GetTotal()
	if err != nil {
		logs.Error("ModUser.GetTotal", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	user.ReturnList(ctx, user_list, user_total)
}

func (user *UserController) AddUser(ctx *gin.Context) {

	userStruct := models.User{}

	if err := ctx.ShouldBindBodyWith(&userStruct, binding.JSON); err != nil {
		logs.Error("ShouldBindBodyWith", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}
	user_id, err := models.ModUser.AddOne(userStruct)
	if err != nil {
		logs.Error("ModUser.AddOne error:", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}
	logs.Debug("add user", user_id, userStruct)
	user.ReturnSucc(ctx, user_id)
}

func (user *UserController) DelUser(ctx *gin.Context) {
	user_id_str := ctx.Param("id")
	user_id, _ := strconv.ParseInt(user_id_str, 10, 64)

	err := models.ModUser.DeleteOne(user_id)

	if err != nil {
		logs.Error("ModUser.DeleteOne error:", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	logs.Debug("del user", user_id)

	user.ReturnSucc(ctx, "ok")
}

func (user *UserController) SetUser(ctx *gin.Context) {

	user_id_str := ctx.Param("id")
	user_id, _ := strconv.ParseInt(user_id_str, 10, 64)

	userStruct := models.User{}

	if err := ctx.ShouldBindBodyWith(&userStruct, binding.JSON); err != nil {
		logs.Error("ShouldBindBodyWith", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	if err := models.ModUser.UpdateOne(user_id, &userStruct); err != nil {
		logs.Error("ModUser.UpdateOne error", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	user.ReturnSucc(ctx, "ok")
}

func (user *UserController) GetUser(ctx *gin.Context) {

	user_id_str := ctx.Param("id")
	user_id, _ := strconv.ParseInt(user_id_str, 10, 64)

	userInfo, err := models.ModUser.GetOne(user_id)
	if err != nil {
		logs.Error("ModUser.GetOne error:", err.Error())
		user.ReturnFail(ctx, err.Error())
		return
	}

	user.ReturnSucc(ctx, userInfo)
}
