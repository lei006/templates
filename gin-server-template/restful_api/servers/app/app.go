package app

import (
	"fmt"
	"gin-server-template/servers/setting"

	"github.com/gin-gonic/gin"
)

func Run(engine *gin.Engine) {
	engine.Run(fmt.Sprintf(":%d", setting.HttpPort))
}
