// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"embed"
	"gin-server-template/middlewares"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func engine_read_file(engine *gin.Engine, views embed.FS, relativePath string, filePath string) {

	engine.GET(relativePath, func(c *gin.Context) {

		file_data, err := views.ReadFile(filePath)
		if err != nil {
			c.String(http.StatusNotFound, "embed : no find file:"+filePath)
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", file_data)
		}
	})
}

func engine_read_dir(engine *gin.Engine, views embed.FS, relativePath string, dirPath string) {
	engine.StaticFS(relativePath, AssetHandler(relativePath, views, dirPath))
}

func Init(engine *gin.Engine, views embed.FS, views_static embed.FS) {

	engine.Use(middlewares.Cors()) //处理跨域
	engine.Use(gin.Recovery())

	//engine.StaticFS("/store", http.Dir(setting.StoreDataPath))

	///////////////////////////////////////////////
	// 映射嵌入式静态文件
	engine_read_dir(engine, views, "/static/", "views/static/")
	engine_read_file(engine, views, "/index.html", "views/index.html")
	engine_read_file(engine, views, "/favicon.ico", "views/favicon.ico")
	engine_read_file(engine, views, "/", "views/index.html")

	//engine.StaticFS("/",http.FS(views))

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

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// AssetHandler returns an http.Handler that will serve files from
// the Assets embed.FS. When locating a file, it will strip the given
// prefix from the request and prepend the root to the filesystem.
func AssetHandler(prefix string, assets embed.FS, root string) http.FileSystem {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		// If we can't find the asset, fs can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}

		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	})

	return http.FS(handler)
}
