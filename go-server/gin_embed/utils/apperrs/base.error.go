package apperrs

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code        int
	Description string
}

func (h *AppError) Error() string {
	return fmt.Sprintf("status code %d: %s", h.Code, h.Description)
}

func MakeAppError(code int, desc string) AppError {
	return AppError{code, desc}
}

var ERR_SERVER_INTERNAL = errors.New("服务器内部错误")
var ERR_SERVER_MYSQL_ERROR = MakeAppError(50010, "数据库内部错误")
var ERR_NO_FIND_ERROR = MakeAppError(50010, "未找到指定数据")
var ERR_CREATE_NO_FIND_ERROR = MakeAppError(50010, "创建却未找到数据")
var ERR_CREATE_NEW_DATA_ERROR = MakeAppError(50010, "创建新的数据失败")
var ERR_NO_FIND_TOKEN_ERROR = MakeAppError(50010, "未找到TOKEN信息数据")
var ERR_PARAME_ERROR = MakeAppError(50010, "参数错误")
var ERR_TOKEN_TIMEOUT_ERROR = MakeAppError(50014, "登录已过期")

var ERR_NO_FIND_USER_ERROR = MakeAppError(50010, "未找到用户")
var ERR_USER_LOGIN_ERROR = MakeAppError(50010, "用户名密码错误")
