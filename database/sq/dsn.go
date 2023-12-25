package sq

import (
	"fmt"
	"goplay/config"
	"time"

	driver "github.com/go-sql-driver/mysql"
)

// Create db first:
// CREATE DATABASE IF NOT EXISTS gormdb
// CHARACTER SET utf8
// COLLATE utf8_general_ci
// WARNING:
// Although you can save time.Time by setting Loc: time.UTC and ParseTime: true here,
// you cannot specify the format used for JSON output.
// It's better to define your own time types.
func buildDSN(c config.Conn, dbName string) *driver.Config {
	return &driver.Config{
		User:   c.User,
		Passwd: c.Pass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", c.Host, c.Port),
		DBName: dbName,
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
		Loc:                  time.UTC,
		AllowNativePasswords: true,
		ParseTime:            true, // Enable parsing time string to golang time.Time
	}
}
