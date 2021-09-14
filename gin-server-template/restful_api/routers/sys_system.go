package routers

import (
	"gin-server-template/controllers"

	"github.com/gin-gonic/gin"
)

func initSystem(pub_group *gin.RouterGroup, check_group *gin.RouterGroup) {

	system := controllers.SystemController{}

	pub_group.GET("/system/about", system.About)
	check_group.GET("/system/license", system.GetLicense)
	check_group.PUT("/system/license", system.SetLicense)
	check_group.POST("/system/restart", system.Restart)
}
