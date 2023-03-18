// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dicom3d/assets"
	"dicom3d/controllers"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func router_embed_assets(engine *gin.Engine) {

	///////////////////////////////////////////////
	// 嵌入式，映射静态文件: 可以保持目录结构
	Static, _ := fs.Sub(assets.Assets, "assets")
	engine.Any("/assets/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(Static))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	engine.Handle("GET", "/index.html", func(ctx *gin.Context) {
		// 注意 assets/index.htm 文件，扩展名，绝不可以为 html
		ctx.FileFromFS("assets/index.htm", http.FS(assets.Assets))
	})
	engine.Handle("GET", "/vite.svg", func(ctx *gin.Context) {
		ctx.FileFromFS("assets/vite.svg", http.FS(assets.Assets))
	})

	/*
		engine.Any("/static/*filepath", func(c *gin.Context) {
			staticServer := http.FileServer(http.FS(assets.Static))
			staticServer.ServeHTTP(c.Writer, c.Request)
		})

		engine.Handle("GET", "/index.html", func(ctx *gin.Context) {
			ctx.FileFromFS("views/index.htm", http.FS(assets.Views))
		})
		engine.Handle("GET", "/favicon.ico", func(ctx *gin.Context) {
			ctx.FileFromFS("views/favicon.ico", http.FS(assets.Views))
		})

	*/

	/*
		///////////////////////////////////////////////
		// 非嵌入式，映射静态文件
		engine.StaticFS("/static", http.Dir(config.WorkPath+"views/static"))
		engine.StaticFile("/index.html", config.WorkPath+"views/index.html")
		engine.StaticFile("/favicon.ico", config.WorkPath+"views/favicon.ico")
		engine.StaticFile("/", config.WorkPath+"views/index.html")
	*/

}

func Init(engine *gin.Engine) {
	engine.Use(gin.Recovery())
	//engine.Use(middlewares.Error()) //统一出错处理
	//engine.Use(middlewares.Cors())  //处理跨域

	// 路由嵌入的资源
	router_embed_assets(engine)

	checkGroup := engine.Group("/api") //验证jwt

	initAbout(checkGroup)
}

func initAbout(router *gin.RouterGroup) {

	about_controller := controllers.AboutController{}

	router.GET("/about", about_controller.About)
	router.GET("/license", about_controller.GetLicense)
	router.GET("/license22", about_controller.GetLicense)
}
