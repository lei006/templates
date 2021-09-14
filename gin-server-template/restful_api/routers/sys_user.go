package routers

import (
	"gin-server-template/controllers"

	"github.com/gin-gonic/gin"
)

func initUser(pub_group *gin.RouterGroup, check_group *gin.RouterGroup) {

	user := controllers.UserController{}

	check_group.GET("/users", user.GetList)
	check_group.POST("/user", user.AddUser)
	check_group.DELETE("/user/:id", user.DelUser)
	check_group.PUT("/user/:id", user.SetUser)
	check_group.GET("/user/:id", user.GetUser)

}
