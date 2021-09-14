// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gin-server-template/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {

	engine.Use(middlewares.Cors()) //处理跨域
	engine.Use(gin.Recovery())

	///////////////////////////////////////////////
	// 映射静态文件
	//engine.StaticFS("/store", http.Dir(setting.StoreDataPath))
	engine.StaticFS("/static", http.Dir("views/static"))
	engine.StaticFile("/index.html", "views/index.html")
	engine.StaticFile("/favicon.ico", "views/favicon.ico")
	engine.StaticFile("/", "views/index.html")

	///////////////////////////////////////////////
	// 创建公共路由与JWT路由
	pubGroup := engine.Group("/api")   //不验证jwt
	checkGroup := engine.Group("/api") //验证jwt

	//checkGroup.Use(middlewares.JWTAuth()) //加载 jwt中间键
	checkGroup.Use(middlewares.TokenAuth()).Use(middlewares.AuthLimit()) //加载 token 检查, token 权限

	///////////////////////////////////////////////
	// 各个模块 api 路由
	initAuth(pubGroup, checkGroup)   //注册认证部份
	initUser(pubGroup, checkGroup)   //注册用户部份
	initSystem(pubGroup, checkGroup) //注册系统部份
	initSetup(pubGroup, checkGroup)  //注册配置部份

}
