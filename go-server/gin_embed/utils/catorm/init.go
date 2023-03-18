package catorm

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func InitMySql(dataSource string, maxIdleConns int, maxOpenConns int, maxLifetime time.Duration) error {
	return getMysqlDB().ConnectMySQL(dataSource, maxIdleConns, maxOpenConns, maxLifetime)
}

func NewDB() *gorm.DB {
	return getMysqlDB().NewDB()
}

func RegisterModel(data interface{}) error {
	return getMysqlDB().RegisterModel(data)
}

type Model struct {
	Id        int64     `gorm:"primarykey,auto_increment" bson:"id" json:"id"`
	CreatedAt time.Time `gorm:"index;type(datetime)" bson:"created_at" json:"created_at"` //创建时间
	UpdatedAt time.Time `gorm:"index;type(datetime)" bson:"updated_at" json:"updated_at"` //修改时间
}
