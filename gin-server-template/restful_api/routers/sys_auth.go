package routers

import (
	"gin-server-template/controllers"

	"github.com/gin-gonic/gin"
)

func initAuth(pub_group *gin.RouterGroup, check_group *gin.RouterGroup) {
	auth := controllers.AuthController{}
	pub_group.POST("/auth/login", auth.Login)
	check_group.POST("/auth/logout", auth.Logout)
	check_group.GET("/auth/info", auth.Info)
}
