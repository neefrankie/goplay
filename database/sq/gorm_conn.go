package sq

import (
	"goplay/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQL(c config.Conn) (*gorm.DB, error) {

	return gorm.Open(mysql.New(mysql.Config{
		DSNConfig: buildDSN(c, "gormdb"),
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

func ConnectSQLite() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

func MustConnectSQLite() *gorm.DB {
	db, err := ConnectSQLite()
	if err != nil {
		panic(err)
	}

	return db
}
