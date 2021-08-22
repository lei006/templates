package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func PathMustExist(fileName string) error {

	_, err := os.Stat(fileName)
	if os.IsNotExist(err) == false {
		return nil
	}

	return os.MkdirAll(fileName, os.ModePerm) //创建多级目录
}

func GetPaths(dirPth string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			files = append(files, fi.Name())
			continue
		}
	}

	return files, nil
}

func GetFiles(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	//PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			//files = append(files, dirPth+PthSep+fi.Name())
			files = append(files, fi.Name())
		}
	}

	return files, nil
}
