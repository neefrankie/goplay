package main

import (
	driver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func buildDSN() string {
	cfg := &driver.Config{
		User:   "sampadm",
		Passwd: "secret",
		DBName: "gormdb",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		// Always use UTC time.
		// Pay attention to how string values are specified.
		// The string value provided to MySQL must be quoted in single quote for this driver to work,
		// which means the single quote itself must be included in the string value.
		// The resulting string value passed to MySQL should look like: `%27<you string value>%27`
		// See ASCII Encoding Reference https://www.w3schools.com/tags/ref_urlencode.asp
		Params: map[string]string{
			"time_zone": `'+00:00'`,
		},
		Collation:            "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
	}

	return cfg.FormatDSN()
}

func connect(dsn string) (*gorm.DB, error) {

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
