package controllers

import (
	"gin-server-template/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SetupController struct {
	baseController
}

type SetupItem struct {
	Name string `bson:"name" json:"name"`
	Data string `bson:"data" json:"data"`
	Desc string `bson:"desc" json:"desc"`
}

func (controller *SetupController) GetList(ctx *gin.Context) {

	//users := []User{{ID: 123, Name: "张三1", NickName: "11"}, {ID: 456, Name: "李四"}}
	object_list, err := models.ModSetup.GetAll()
	if err != nil {
		logs.Error("ModSetup.GetAll", err.Error())
		controller.ReturnFail(ctx, err.Error())
		return
	}

	controller.ReturnList(ctx, object_list, int64(len(object_list)))
}

func (controller *SetupController) SetOne(ctx *gin.Context) {

	dataField := SetupItem{}
	if err := ctx.ShouldBindBodyWith(&dataField, binding.JSON); err != nil {

		logs.Error("ShouldBindBodyWith ", err.Error())
		controller.ReturnFail(ctx, err.Error())
		return
	}

	err := models.ModSetup.SetString(dataField.Name, dataField.Data, dataField.Desc)
	if err != nil {
		logs.Error("PresetController  Update dataField error ", err.Error())
		controller.ReturnFail(ctx, err.Error())
		return
	}

	controller.ReturnSucc(ctx, "ok")
}

func (controller *SetupController) GetOne(ctx *gin.Context) {

	name := ctx.Param("id")

	if name == "thumbnail_size" {

	}

	info, err := models.ModSetup.GetString(name)
	if err != nil {
		logs.Error("ModSetup.GetOne error:", err.Error())
		controller.ReturnFail(ctx, err.Error())
		return
	}

	controller.ReturnSucc(ctx, info)
}
