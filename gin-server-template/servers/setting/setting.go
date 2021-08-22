package setting

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
)
