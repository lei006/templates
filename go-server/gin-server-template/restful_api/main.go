package main

import (
	"embed"
	"gin-server-template/services/app"
	"gin-server-template/services/setting"
	"gin-server-template/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/lei006/go-assist/tools/daemontool"
)

//go:embed views
var views embed.FS

//go:embed views/static/*
var views_static embed.FS

func main() {

	if setting.RunMode == setting.RunMode_PROD || utils.Is_RunAtVs() == false {
		daemontool.DefDaemonTool.Run(setting.AppName, setting.AppDesc, app_run)
	} else {
		app_run()
	}
}

func app_run() {

	logs.Notice("app run at port: ", setting.HttpPort)

	app.Run(gin.Default(), views, views_static)

}



