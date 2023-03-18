package catorm

import (
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormMysql struct {
	db *gorm.DB
}

// GormMysql 初始化Mysql数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (orm *gormMysql) ConnectMySQL(dataSource string, maxIdleConns int, maxOpenConns int, maxLifetime time.Duration) error {

	mysqlConfig := mysql.Config{
		DSN:                       dataSource, // DSN data source name
		DefaultStringSize:         512,        // string 类型字段的默认长度
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(maxLifetime)

	orm.db = db

	return nil
}

func (orm *gormMysql) Config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}

func (orm *gormMysql) NewDB() *gorm.DB {
	return orm.db
}

func (orm *gormMysql) RegisterModel(data interface{}) error {
	return orm.db.AutoMigrate(data)
}

var ins_mysql *gormMysql

var once sync.Once

func getMysqlDB() *gormMysql {
	once.Do(func() {
		ins_mysql = &gormMysql{}
	})
	return ins_mysql
}
