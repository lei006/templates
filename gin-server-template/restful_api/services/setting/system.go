package setting

import (
	"flag"
	"gin-server-template/utils"
	"os"
	"strings"
	"time"

	"github.com/beego/beego/v2/adapter/config"
	"github.com/beego/beego/v2/core/logs"
)

const (
	RunMode_DEV  = "dev"
	RunMode_PROD = "prod"
)

var (
	AppName   = "gin-server-template"
	AppDesc   = "服务器模板"
	AppIssuer = "徐州云创电子科技有限公司"
	HttpPort  = int(8180)
	BuildTime string

	RunMode = RunMode_PROD

	WorkPath   = ""
	log_enable = "true"

	HardSn string //硬件id

	TokenKeepTime = int(365 * 24 * 60 * 60) //秒
)

func init() {
	BuildTime = utils.BuildTime()
	//////////////////////////////
	// 加载硬件id
	hardSn, err := utils.GetPhysicalID()
	if err != nil {
		logs.Error("server exit : get hard sn error =", err.Error())
		os.Exit(-1)
	}
	HardSn = hardSn

	logs.Debug("utils.GetWorkPath() ", utils.GetBinPath())

	if utils.Is_RunAtVs() {
		//运行在vs 中
		WorkPath = utils.GetWorkPath() + "/"
	} else {
		WorkPath = utils.GetBinPath() + "/"
	}

	//加载配置文件
	err = ReLoadSysConfig()
	if err != nil {
		logs.Error("server exit : read config file error =", err.Error())
		time.Sleep(time.Second)
		os.Exit(-1)
	}

	//////////////////////////////////////////////////////////
	// 配置运行路径 -- 不在 vs
	if RunMode == RunMode_PROD && !utils.Is_RunAtVs() {
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置日志...
	if log_enable == "true" {
		log_filename := WorkPath + strings.ToLower(AppName) + ".log"
		logs.SetLogger(logs.AdapterFile, `{"filename":"`+log_filename+`"}`)
	}

	logs.Debug("")
	logs.Debug("")
	logs.Debug("----------------------------------------------------------------------")
	logs.Debug("----------------------------------------------------------------------")

	logs.Debug("RunMode =", RunMode)
	logs.Debug("BuildTime =", BuildTime)
	logs.Debug("HardSn =", HardSn)

	logs.Warn("webdcm run at path: ", WorkPath)

}

func ReLoadSysConfig() error {

	configpath := flag.String("c", "conf/app.conf", "The default to load the conf/app.conf")
	flag.Parse()

	p_config_port := flag.Int("p", HttpPort, "defaut")

	config_filenamepath := WorkPath + *configpath
	//config_filenamepath := "D:/Work/webdcm/viewer/trunk_gin/" + *configpath
	//config_filenamepath := "D:/Work/webdcm/viewer/trunk_gin/conf/app.conf"

	config_filenamepath = strings.Replace(config_filenamepath, "/", "\\", -1) //将/替换成\\

	logs.Notice("read config file :", config_filenamepath)

	iniconf, err := config.NewConfig("ini", config_filenamepath)
	//iniconf, err := config.NewConfig("ini", "D:\\Work\\webdcm\\viewer\\trunk_gin\\conf\\app.conf")
	if err != nil {
		return err
	}

	HttpPort = iniconf.DefaultInt("httpport", *p_config_port)
	RunMode = iniconf.DefaultString("runmode", RunMode)

	return nil
}
