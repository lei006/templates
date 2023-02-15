package main

import (
	"fmt"
	"demo_embed_mini/config"
	"demo_embed_mini/routers"
	"demo_embed_mini/utils"

	"github.com/gin-gonic/gin"
	"github.com/beego/beego/v2/core/logs"
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

	////////////////////////////////////////
	// 4. 设置缺省项
	app_init_system()


	logs.Notice("app run at port: ", config.HttpPort)
	engine.Run(fmt.Sprintf(":%d", config.HttpPort))

}

func app_init_system() {

}
