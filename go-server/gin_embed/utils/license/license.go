package license

import (
	"errors"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/lei006/go-assist/license"
)

func CheckLicenseFile(filename string, appName string, Issuer, HardSn string) (string, *license.LicenseClaims, error) {

	// 1. 读文件
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", nil, err
	}

	lic_data := string(content)

	// 2. 检查授权数据
	claims, err := CheckLicenseData(lic_data, appName, Issuer, HardSn)

	return lic_data, claims, err

}

func CheckLicenseData(lic_data string, AppName string, Issuer, hardSn string) (*license.LicenseClaims, error) {
	if lic_data == "" {
		return nil, errors.New("注册信息为空")
	}

	// 1. 检查授权码是否可行 ...
	dec_data, err := license.Decryption(lic_data)
	if err != nil {
		logs.Error("注册码无效:", err.Error())
		return nil, errors.New("注册码无效")
	}

	if dec_data.Hardsn != hardSn {
		logs.Error("注册码无效: 硬件码不同")
		return nil, errors.New("注册码无效: 硬件码不同:" + hardSn + "-" + dec_data.Hardsn)
	}
	if dec_data.AppName != AppName {
		logs.Error("注册码无效: 程序名不同")
		return nil, errors.New("注册码无效: 程序名不同:" + dec_data.AppName + "-" + AppName)
	}
	if dec_data.Issuer != "徐州云宝电子科技" {
		logs.Error("注册码无效: 发行人错误")
		return nil, errors.New("注册码无效: 发行人错误:" + dec_data.Issuer)
	}

	return dec_data, nil
}
