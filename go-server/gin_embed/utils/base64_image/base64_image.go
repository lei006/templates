package base64_image

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
)

func WriteFile(filenamePath string, base64_image_content string) (string, error) {

	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64_image_content)
	if !b {
		return "", errors.New("base64_image_content not is image")
	}

	re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	allData := re.FindAllSubmatch([]byte(base64_image_content), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取

	base64Str := re.ReplaceAllString(base64_image_content, "")

	var file string = filenamePath + "." + fileType
	byte, _ := base64.StdEncoding.DecodeString(base64Str)

	err := ioutil.WriteFile(file, byte, 0666)
	if err != nil {
		return "", err
	}

	return file, nil
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

}

func ReadFile(filenamePath string) (string, error) {
	ff, err := ioutil.ReadFile(filenamePath) //我还是喜欢用这个快速读文件
	if err != nil {
		return "", err
	}
	out_base64 := base64.StdEncoding.EncodeToString(ff) // 文件转base64
	return "data:image/jpeg;base64," + out_base64, nil
}
