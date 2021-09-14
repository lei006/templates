package app

import (
	"errors"
	"fmt"
	"gin-server-template/services/setting"
)

func InitDB(path string) error {

	driverName := setting.DbDriverName
	dbDataSource := setting.DbDataSource

	fmt.Println("InitDB DriverName", driverName)
	fmt.Println("InitDB DataSource", dbDataSource)

	// beego.AppConfig.DefaultString("DbDataSource", path+Default_DbDataSource)
	var err error

	if driverName == "sqlite3" {
		if err = InitSqlite(dbDataSource); err != nil {
			return err
		}
	} else if driverName == "mysql" {
		if err = InitMySql(dbDataSource); err != nil {
			return err
		}
	} else {
		return errors.New("暂未支持数据库:" + driverName)
	}

	return nil
}
