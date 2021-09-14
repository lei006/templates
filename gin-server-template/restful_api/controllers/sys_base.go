package controllers

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

type baseController struct {
}

type JsonReturn struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string      `json:"message"`
	Now  int64       `json:"now"`
}

func (this *JsonReturn) ToJson() (string, error) {

	b, err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (base *baseController) ApiJsonReturn(ctx *gin.Context, code int, msg string, data interface{}) {

	origin := ctx.Request.Header.Get("Origin")
	ctx.Header("Access-Control-Allow-Origin", origin)
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

	var _jsonReturn JsonReturn
	_jsonReturn.Msg = msg
	_jsonReturn.Code = code
	_jsonReturn.Data = data
	_jsonReturn.Now = time.Now().Unix()
	//c.Data["json"] = _jsonReturn //将结构体数组根据tag解析为json
	//c.ServeJSON()                //对json进行序列化输出
	//c.StopRun()                  //终止执行逻辑
	ctx.JSON(200, _jsonReturn)
	ctx.Abort()
}

type JsonDataList struct {
	Items interface{} `json:"items"` //Data字段需要设置为interface类型以便接收任意数据
	Total int64       `json:"total"`
}

type JsonReturnList struct {
	Code int          `json:"code"`
	Data JsonDataList `json:"data"` //Data字段需要设置为interface类型以便接收任意数据
	Msg  string       `json:"message"`
	Now  int64        `json:"now"`
}

func (c *baseController) ReturnList(ctx *gin.Context, data interface{}, total int64) {

	/*
		origin := ctx.Request.Header.Get("Origin")
		ctx.Header("Access-Control-Allow-Origin", origin)
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	*/
	var _ret JsonReturnList
	_ret.Msg = "success"
	_ret.Code = 20000
	_ret.Data.Items = data
	_ret.Data.Total = total
	_ret.Now = time.Now().Unix()

	//c.Data["json"] = _ret //将结构体数组根据tag解析为json
	//c.ServeJSON()         //对json进行序列化输出
	//c.StopRun()           //终止执行逻辑
	ctx.JSON(200, _ret)
	ctx.Abort()
}

func (base *baseController) ReturnSucc(ctx *gin.Context, data interface{}) {
	base.ApiJsonReturn(ctx, 20000, "success", data)
}

func (base *baseController) ReturnFail(ctx *gin.Context, msg string) {
	base.ApiJsonReturn(ctx, 50010, msg, nil)
}

func (base *baseController) ReturnFailServer(ctx *gin.Context, msg string) {

	base.ApiJsonReturn(ctx, 50000, "server error: "+msg, nil)
}

func (base *baseController) ReturnFailCode(ctx *gin.Context, code int, msg string) {
	base.ApiJsonReturn(ctx, code, msg, nil)
}
