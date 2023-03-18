package controllers

import (
	"dicom3d/config"
	"dicom3d/utils/license"

	"github.com/gin-gonic/gin"
)

type AboutController struct {
	BaseController
}

type AboutDataItem struct {
	Key string      `bson:"key" json:"key"`
	Val interface{} `bson:"val" json:"val"`
}

func (system *AboutController) About(ctx *gin.Context) {

	list := []AboutDataItem{}
	list = append(list, AboutDataItem{Key: "程序名", Val: config.AppName})
	list = append(list, AboutDataItem{Key: "编译时间", Val: config.BuildTime})
	list = append(list, AboutDataItem{Key: "硬件码", Val: config.HardSn})
	//list = append(list, types.DataItem{Key: "版权所有", Val: this.Owner})
	Lic_claims, err := license.CheckLicenseData(config.LicenseData, config.AppName, config.AppIssuer, config.HardSn)
	if err == nil {
		list = append(list, AboutDataItem{Key: "是否授权", Val: "已授权"})
		list = append(list, AboutDataItem{Key: "授权人", Val: Lic_claims.Issuer})
		list = append(list, AboutDataItem{Key: "授权标题", Val: Lic_claims.Subject})
	} else {
		list = append(list, AboutDataItem{Key: "是否授权", Val: "未授权"})
	}

	system.Success(ctx, list)
}

func (system *AboutController) GetLicense(ctx *gin.Context) {

	Lic_claims, err := license.CheckLicenseData(config.LicenseData, config.AppName, config.AppIssuer, config.HardSn)
	if err != nil {
		system.FailCode(ctx, 401, err.Error())
		return
	}

	system.Success(ctx, Lic_claims)
}
