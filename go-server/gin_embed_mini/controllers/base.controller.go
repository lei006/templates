package controllers

import (
	"encoding/base64"
	"encoding/json"
	"demo_embed_mini/utils/apperrs"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (base *BaseController) cors(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
}

func (base *BaseController) ImageReturn(ctx *gin.Context, imageBase64 string) {

	base.cors(ctx)
	imageBase64 = strings.Replace(imageBase64, "data:image/jpeg;base64,", "", 1)
	ddd, err := base64.StdEncoding.DecodeString(imageBase64) //成图片文件并把文件写入到buffer
	if err != nil {
		return
	}
	ctx.Writer.Header().Set("Content-Type", "image/jpeg")
	ctx.Writer.Write(ddd)
}

func (base *BaseController) JsonReturn(ctx *gin.Context, code int, msg string, data interface{}) {

	base.cors(ctx)

	var _jsonReturn JsonReturn
	_jsonReturn.Msg = msg
	_jsonReturn.Code = code
	_jsonReturn.Data = data
	_jsonReturn.Now = time.Now()
	ctx.JSON(200, _jsonReturn)
	ctx.Abort()
}

func (base *BaseController) JsonReturnList(ctx *gin.Context, data interface{}, total int64) {

	base.cors(ctx)

	var _ret JsonReturnList
	_ret.Msg = "success"
	_ret.Code = 20000
	_ret.Data.Items = data
	_ret.Data.Total = total
	_ret.Now = time.Now()

	ctx.JSON(200, _ret)
	ctx.Abort()
}

func (base *BaseController) Success(ctx *gin.Context, data interface{}) {
	base.JsonReturn(ctx, 20000, "success", data)
}
func (base *BaseController) SuccessOk(ctx *gin.Context) {
	base.JsonReturn(ctx, 20000, "success", "ok")
}

func (base *BaseController) FailErr(ctx *gin.Context, err apperrs.AppError) {
	base.JsonReturn(ctx, err.Code, err.Description, nil)
}

func (base *BaseController) Fail(ctx *gin.Context, msg string) {
	base.JsonReturn(ctx, 50010, msg, nil)
}

func (base *BaseController) FailMessage(ctx *gin.Context, msg string) {
	base.JsonReturn(ctx, 50013, msg, nil)
}

func (base *BaseController) FailCode(ctx *gin.Context, code int, msg string) {
	base.JsonReturn(ctx, code, msg, nil)
}

func (base *BaseController) GetUserId(ctx *gin.Context) int64 {
	userid_ctx, is_exist := ctx.Get("user_id")
	if is_exist == false { //不存在
		base.FailErr(ctx, apperrs.ERR_PARAME_ERROR)
		return 0
	}
	login_user_id := userid_ctx.(int64)

	return login_user_id
}

func (base *BaseController) SetUserId(ctx *gin.Context, user_id int64) {
	ctx.Set("user_id", user_id)
}

func (base *BaseController) GetParamInt64(ctx *gin.Context, param_name string) int64 {
	param_data_str := ctx.Param(param_name)
	param_data, err := strconv.ParseInt(param_data_str, 10, 64)
	if err != nil {
		logs.Error("参数出错:" + param_name + param_data_str)
		base.FailErr(ctx, apperrs.ERR_PARAME_ERROR)
		return -1
	}

	return param_data
}

func (base *BaseController) GetQueryInt64(ctx *gin.Context, param_name string) int64 {
	param_data_str := ctx.Query(param_name)
	param_data, err := strconv.ParseInt(param_data_str, 10, 64)
	if err != nil {
		logs.Error("参数出错:" + param_name + param_data_str)
		base.FailErr(ctx, apperrs.ERR_PARAME_ERROR)
		return -1
	}

	return param_data
}

type JsonReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string      `json:"message"`
	Now  time.Time   `json:"now"`
}

func (this *JsonReturn) ToJson() (string, error) {

	b, err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

type JsonDataList struct {
	Items interface{} `json:"items"` //Data字段需要设置为interface类型以便接收任意数据
	Total int64       `json:"total"`
}

type JsonReturnList struct {
	Code int          `json:"code"`
	Data JsonDataList `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string       `json:"message"`
	Now  time.Time    `json:"now"`
}
