package orm

import (
	"goplay/web/config"

	"goplay/web/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQL(c config.Conn) (*gorm.DB, error) {

	return gorm.Open(mysql.New(mysql.Config{
		DSNConfig: db.BuildDSN(c, "gormdb"),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func MustNewMySQL(c config.Conn) *gorm.DB {
	myDB, err := NewMySQL(c)

	if err != nil {
		panic(err)
	}

	return myDB
}

func getMyDB() *gorm.DB {
	return MustNewMySQL(config.MustLoad().GetMySQLConn())
}
