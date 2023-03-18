package config

import (
	"dicom3d/utils"
	"dicom3d/utils/license"
	"flag"
	"os"

	"strings"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	"github.com/lei006/go-assist/hardsn"
)

const (
	RunMode_DEV  = "dev"
	RunMode_PROD = "prod"

	STORE_DIR_NAME = "store"
	CVR_DIR_NAME   = "cvr"

	PUBLISH_KEY = `-----BEGIN ecc public key-----
	MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQBxEs4MdO861+CGAkqOA2RqYRyHsxr
	lwVw0NBhy8ipXQFqa39mslJiLDVs5Tipl/mVXmJrBnKveHL/DJPDPfw6QEsBs7/Y
	Xu3blP8l22hPEDyW4bZICgKnbiArwi16xvnQWXDk+5FCWS4A9NskefeENle6Tiqx
	l4fPkd2/ML63yYfFkTs=
	-----END ecc public key-----`
)

var (
	AppName   = "Dicom3D"
	AppDesc   = "Web三维工作站-服务器"
	AppIssuer = "徐州云创电子有限公司"
	HttpPort  = int(8180)
	BuildTime string
	Debug     = false
	HardSn    string //硬件id

	TokenKeepTime = int(60) //分钟

	TokenSigningKey = "abcde" //分钟

	WorkPath = "./"

	log_enable = "true"

	license_file = "license.lic"
	LicenseData  = "" //授权数据
)

func init() {

	BuildTime = utils.BuildTime()

	//////////////////////////////
	// 加载硬件id
	id_md5, err := hardsn.GetPhysicalID()
	if err != nil {
		logs.Error("server exit : get hard sn error =", err.Error())
		os.Exit(-1)
	}

	HardSn = id_md5

	if utils.Is_RunAtVs() {
		//运行在vs 中
		WorkPath = utils.GetWorkPath() + "/"
	} else {
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置运行路径 -- 不在 vs
	if Debug == false && !utils.Is_RunAtVs() {
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置日志...
	log_filename := WorkPath + strings.ToLower(AppName) + ".log"
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+log_filename+`", "level":7}`)

	logs.Debug("WorkPath = " + WorkPath)
	logs.Debug("BinPath = " + utils.GetBinPath())
	logs.Debug("GetWorkPath = " + utils.GetWorkPath())
	logs.Debug("Is_RunAtVs = ", utils.Is_RunAtVs())

	//加载配置文件
	err = ReLoadConfig()
	if err != nil {
		logs.Error("server exit : read config file error =", err.Error())
		os.Exit(-1)
	}

	if WorkPath == "" {
		//配置文件中没有配置，则用默认的..
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置日志...
	if log_enable == "true" {
		log_filename := WorkPath + strings.ToLower(AppName) + ".log"
		logs.SetLogger(logs.AdapterFile, `{"filename":"`+log_filename+`", "level":7}`)
	}

	//////////////////////////////////////////////////////////
	// 检查授权文件
	CheckLicenseFile()
}

func CheckLicenseFile() {
	licData, _, err := license.CheckLicenseFile(license_file, AppName, AppIssuer, HardSn)
	if err != nil {
		logs.Error("此服务器未授权:" + err.Error())
	} else {
		//检查OK
		LicenseData = licData
	}
}

func ReLoadConfig() error {

	configpath := flag.String("c", "conf/app.conf", "The default to load the conf/app.conf")
	flag.Parse()

	config_filenamepath := WorkPath + *configpath

	config_filenamepath = strings.Replace(config_filenamepath, "/", "\\", -1) //将/替换成\\

	logs.Notice("read config file :", config_filenamepath)

	iniconf, err := config.NewConfig("ini", config_filenamepath)
	if err != nil {
		return err
	}

	HttpPort = iniconf.DefaultInt("httpport", HttpPort)
	Debug = iniconf.DefaultBool("debug", Debug)
	TokenKeepTime = iniconf.DefaultInt("TokenKeepTime", TokenKeepTime)

	log_enable = iniconf.DefaultString("log_enable", log_enable)

	return nil
}
