package app

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitMongo(dataSource string) error {

	/*
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		orm.RegisterDataBase("default", "sqlite3", dataSource)
		orm.RunSyncdb("default", false, true)
	*/
	fmt.Println(dataSource)
	return nil
}
