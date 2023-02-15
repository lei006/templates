// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"net/http"
	"demo_embed_mini/config"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	//engine.Use(middlewares.Error()) //统一出错处理
	//engine.Use(middlewares.Cors())  //处理跨域

	///////////////////////////////////////////////
	// 映射静态文件
	engine.StaticFS("/static", http.Dir(config.WorkPath+"views/static"))
	engine.StaticFile("/index.html", config.WorkPath+"views/index.html")
	engine.StaticFile("/favicon.ico", config.WorkPath+"views/favicon.ico")
	engine.StaticFile("/", config.WorkPath+"views/index.html")

}
