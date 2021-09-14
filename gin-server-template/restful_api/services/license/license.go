package license

import (
	"fmt"
	"gin-server-template/services/setting"
	"io/ioutil"

	"github.com/lei006/go-assist/tools/licenser_normal"
)

var Licenser *licenser_normal.Licenser

func init() {
	LoadFromeFile(setting.LicenseFile)
}

func LoadFromeFile(filepath string) error {
	Licenser = licenser_normal.MakeLicenser("LiveRTC", setting.AppName, setting.HardSn, setting.PUBLISH_KEY)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = Licenser.SetData(string(data), func(lic_data *licenser_normal.LicenserData) {
	})

	fmt.Println("Licenser ==>", err, string(data))

	return err
}
