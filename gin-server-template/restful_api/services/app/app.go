package app

import (
	"fmt"
	"gin-server-template/routers"
	"gin-server-template/services/setting"

	"github.com/gin-gonic/gin"
)

func _init() error {

	return nil
}

func Run(engine *gin.Engine) {

	_init()

	routers.Init(engine)

	engine.Run(fmt.Sprintf(":%d", setting.HttpPort))
}
