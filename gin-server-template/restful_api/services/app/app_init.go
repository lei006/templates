package app

import (
	"errors"
	"fmt"
	"gin-server-template/services/setting"

	"github.com/beego/beego/v2/adapter/logs"
)

func InitDB(path string) error {

	driverName := setting.DbDriverName
	dbDataSource := setting.DbDataSource

	fmt.Println("InitDB DriverName", driverName)

	// beego.AppConfig.DefaultString("DbDataSource", path+Default_DbDataSource)
	var err error

	if driverName == "sqlite3" {
		dbDataSource = setting.WorkPath + dbDataSource
		logs.Debug("InitSqlite DataSource", dbDataSource)
		if err = InitSqlite(dbDataSource); err != nil {
			return err
		}
	} else if driverName == "mysql" {
		logs.Debug("InitMySql DataSource", dbDataSource)
		if err = InitMySql(dbDataSource); err != nil {
			return err
		}
	} else {
		return errors.New("暂未支持数据库:" + driverName)
	}

	return nil
}
