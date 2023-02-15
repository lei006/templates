package controllers

import (
	"os"
	"demo_embed_mini/config"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SystemController struct {
	BaseController
}

func (system *SystemController) About(ctx *gin.Context) {

	list := []DataItem{}
	list = append(list, DataItem{Key: "程序名", Val: config.AppName})
	list = append(list, DataItem{Key: "编译时间", Val: config.BuildTime})
	list = append(list, DataItem{Key: "硬件码", Val: config.HardSn})
	//list = append(list, types.DataItem{Key: "版权所有", Val: this.Owner})

	system.Success(ctx, list)
}

func (system *SystemController) List(ctx *gin.Context) {
	//ret := server_monitor.GetDiskInfo()
	system_server := servers.SystemServer{}
	item_list, count := system_server.List()

	system.JsonReturnList(ctx, item_list, count)
}

func (system *SystemController) GetOne(ctx *gin.Context) {
	id := system.GetParamInt64(ctx, "id")

	system_server := servers.SystemServer{}
	item, err := system_server.GetOne(id)
	if err != nil {
		system.Fail(ctx, err.Error())
		return
	}

	system.Success(ctx, item)
}

func (system *SystemController) PatchOne(ctx *gin.Context) {

	id := system.GetParamInt64(ctx, "id")

	dataField := DataField{}
	if err := ctx.ShouldBindBodyWith(&dataField, binding.JSON); err != nil {
		logs.Error("ShouldBindBodyWith ", err.Error())
		system.Fail(ctx, err.Error())
		return
	}
	system_server := servers.SystemServer{}
	err := system_server.PatchOneData(id, dataField.Name, dataField.Data)
	if err != nil {
		system.Fail(ctx, err.Error())
		return
	}

	system.Success(ctx, "ok")
}

func (system *SystemController) GetLicense(ctx *gin.Context) {

	system_server := servers.SystemServer{}
	lic, _ := system_server.GetLicense()

	system.Success(ctx, lic)
}

func (system *SystemController) GetAdminPassword(ctx *gin.Context) {

	system_mod := models.SystemModel{}
	item, is_exist := system_mod.GetOneByName("AdminPassword")
	if is_exist == false {
		system.Fail(ctx, "未找到指定项")
		return
	}

	system.Success(ctx, item)
}

func (system *SystemController) Restart(ctx *gin.Context) {

	go func() {
		logs.Notice("server restart")
		os.Exit(-1)
	}()

	system.Success(ctx, "ok")
}
