package utils

import (
	"os"
	"path/filepath"
	"strings"
)

//使用 go run . 运行..
func Is_Go_Run_Mod() bool {
	return GetBinPath() != GetWorkPath()
}

// 运行在 vs中...
func Is_RunAtVs() bool {
	bin_path := GetBinPath()
	return strings.Contains(bin_path, "Temp/go-build")
}

func GetBinPath() string {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func GetWorkPath() string {
	str, _ := os.Getwd()

	return strings.Replace(str, "\\", "/", -1) //将\替换成/
}

/*



func (this *DaemonTool) GetWordPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

*/
