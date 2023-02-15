package reportno

import (
	"errors"
	"strings"
)

var ERR_FORMAT_ERROR = errors.New("格式错误: 未例格式 YC,YYYY,00001 YC,YYMM,00001 YC,YYMMDD,00001 ")

type CodeMaker struct {
	prefix     string //前缀
	dataFormat string //日期长度
	indexLen   int    //序号长度
}

func MakeNewCodeMaker(format string) (*CodeMaker, error) {
	/////////////////////////////////////////////
	// 1. 分析格式
	str_splits := strings.Split(format, ",")
	if len(str_splits) != 3 {
		return nil, ERR_FORMAT_ERROR
	}

	fix0 := str_splits[0]
	fix1 := str_splits[1] //YYYYMMDD //YYMMDD //YYMM //YY
	fix2 := str_splits[2]

	// 3. 根据设定生成 ID
	//date_all := fmt.Sprintf("%d%02d%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	//date_fix := date_all[:min(len(date_all), len(fix2))]
	if fix1 != "YYYY" && fix1 != "YYMM" && fix1 != "YYYYMM" && fix1 != "YYYYMMDD" && fix1 != "YYMMDD" {
		return nil, ERR_FORMAT_ERROR
	}

	index_len := len(fix2)
	if index_len > 18 {
		return nil, ERR_FORMAT_ERROR
	}

	maker := CodeMaker{prefix: fix0,
		dataFormat: fix0,
		indexLen:   index_len,
	}

	return &maker, nil
}

func (maker *CodeMaker) MakeNew(index int64) string {

	return ""
}
