package routers

import (
	"gin-server-template/controllers"

	"github.com/gin-gonic/gin"
)

func initSetup(pub_group *gin.RouterGroup, check_group *gin.RouterGroup) {

	setup := controllers.SetupController{}

	pub_group.GET("/setups", setup.GetList)
	check_group.PUT("/setup", setup.SetOne)
	check_group.PUT("/setup/:id", setup.SetOne)
	check_group.GET("/setup/:id", setup.GetOne)
}
