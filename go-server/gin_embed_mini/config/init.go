package config

import (
	"errors"
	"flag"
	"os"
	"demo_embed_mini/utils"
	"strings"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	"github.com/lei006/go-assist/hardsn"
	"github.com/lei006/go-assist/license"
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
	AppName   = "web-station-server"
	AppDesc   = "Web工作站-服务器"
	AppIssuer = "徐州云创电子科技有限公司"
	HttpPort  = int(8180)
	BuildTime string
	Debug     = false
	HardSn    string //硬件id

	TokenKeepTime = int(60) //分钟

	TokenSigningKey = "abcde" //分钟

	WorkPath = "./"

	log_enable = "true"

	DbDataSource = ""

	LicenseFile = "license.lic"
	LicenseCode = ""

	AdminPassword = "123456"

	CaptureCmd = "ffmpeg.exe -r 20 -probesize 50M -rtbufsize 256M -f dshow -i video=\"Video (00 Pro Capture Mini HDMI)\" -s 1920x1080  -vcodec libx264 -pix_fmt nv12 -r 20 -b:v 2M -maxrate 2M  -bufsize 2M -bf 0 -g 20  -profile:v baseline  -preset veryfast -f rtsp -rtsp_transport tcp"
	CaptureURL = "rtsp://127.0.0.1:554/capture.sdp"
)

func init() {

	//BuildTime = build_time.BuildTime()

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

	AppName = iniconf.DefaultString("appname", AppName)
	AppIssuer = iniconf.DefaultString("issuer", AppIssuer)
	HttpPort = iniconf.DefaultInt("httpport", HttpPort)
	Debug = iniconf.DefaultBool("debug", Debug)
	TokenKeepTime = iniconf.DefaultInt("TokenKeepTime", TokenKeepTime)

	DbDataSource = iniconf.DefaultString("DbDataSource", DbDataSource)
	CaptureCmd = iniconf.DefaultString("CaptureCmd", CaptureCmd)
	CaptureURL = iniconf.DefaultString("CaptureURL", CaptureURL)

	log_enable = iniconf.DefaultString("log_enable", log_enable)

	return nil
}

func DecodeLicense(lic_data string) (*license.LicenseClaims, error) {

	if lic_data == "" {
		return nil, errors.New("没有注册码")
	}

	// 1. 检查授权码是否可行 ...
	dec_data, err := license.Decryption(lic_data)
	if err != nil {
		logs.Error("注册码无效:", err.Error())
		return nil, errors.New("注册码无效")
	}

	if dec_data.Hardsn != HardSn {
		logs.Error("注册码无效: 硬件码不同")
		return nil, errors.New("注册码无效: 硬件码不同")
	}
	if dec_data.AppName != AppName {
		logs.Error("注册码无效: 程序名不同")
		return nil, errors.New("注册码无效: 程序名不同")
	}
	return dec_data, nil
}
