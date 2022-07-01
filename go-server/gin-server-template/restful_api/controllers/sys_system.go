package controllers

import (
	"gin-server-template/entity"
	"gin-server-template/services/license"
	"gin-server-template/services/setting"
	"io/ioutil"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/lei006/go-assist/tools/licenser_normal"
)

type SystemController struct {
	baseController
}

func (system *SystemController) About(ctx *gin.Context) {

	list := []entity.DataItem{}
	list = append(list, entity.DataItem{Key: "程序名", Val: setting.AppName})
	list = append(list, entity.DataItem{Key: "编译时间", Val: setting.BuildTime})
	list = append(list, entity.DataItem{Key: "硬件码", Val: setting.HardSn})
	//list = append(list, types.DataItem{Key: "版权所有", Val: this.Owner})

	system.ReturnSucc(ctx, list)
}

func (system *SystemController) GetLicense(ctx *gin.Context) {
	list := license.Licenser.GetInfo()
	system.ReturnSucc(ctx, list)
}

func (system *SystemController) SetLicense(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logs.Error("SetLicense Read Request Body error:", err.Error())
		system.ReturnFail(ctx, err.Error())
		return
	}

	err = license.Licenser.SetData(string(body), func(lic_data *licenser_normal.LicenserData) {
		license_str, err := lic_data.ToJson()
		if err != nil {
			logs.Error("license data to json error: ", err.Error())
			return
		}
		err = ioutil.WriteFile(setting.LicenseFile, []byte(license_str), 0666) //写入文件(字节数组)
		if err != nil {
			logs.Error("write license file error: ", err.Error())
		} else {
			logs.Notice("write license file success")
		}
	})

	if err != nil {
		logs.Error("Licenser LoadData error:", err.Error())
		system.ReturnFail(ctx, err.Error())
		return
	}

	system.ReturnSucc(ctx, "ok")
}

func (system *SystemController) Restart(ctx *gin.Context) {

	go func() {
		logs.Notice("server restart")
		os.Exit(-1)
	}()

	system.ReturnSucc(ctx, "ok")
}
