package app

import (
	"fmt"
	"gin-server-template/routers"
	"gin-server-template/services/setting"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/gin-gonic/gin"
)

func Run(engine *gin.Engine) {

	logs.EnableFuncCallDepth(true)

	////////////////////////////////////
	//1. 实始化数据库
	err := InitDB("")
	if err != nil {
		fmt.Println("InitDB error:", err.Error())
		return
	}

	routers.Init(engine)

	engine.Run(fmt.Sprintf(":%d", setting.HttpPort))
}
