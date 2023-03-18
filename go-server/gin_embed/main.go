package main

import (
	"dicom3d/config"
	"dicom3d/routers"
	"dicom3d/utils"
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/lei006/go-assist/tools/daemontool"
)

func main() {
	if config.Debug == false || utils.Is_RunAtVs() == false {
		gin.SetMode(gin.ReleaseMode)
		daemontool.DefDaemonTool.Run(config.AppName, config.AppDesc, app_run)
	} else {
		app_run()
	}
}

func app_run() {

	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号
	logs.EnableFuncCallDepth(true)

	engine := gin.Default()

	////////////////////////////////////
	// 3. 初始化路由
	routers.Init(engine)

	logs.Notice("app run at port: ", config.HttpPort)
	engine.Run(fmt.Sprintf(":%d", config.HttpPort))

}
