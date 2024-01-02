package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Create db first:
// CREATE DATABASE IF NOT EXISTS gormdb
// CHARACTER SET utf8
// COLLATE utf8_general_ci
// WARNING:
// Although you can save time.Time by setting Loc: time.UTC and ParseTime: true here,
// you cannot specify the format used for JSON output.
// It's better to define your own time types.
func BuildDSN(c Conn, dbName string) *mysql.Config {
	return &mysql.Config{
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

func MustGetMyDSN(dbName string) string {
	return BuildDSN(MustLoad().GetMySQLConn(), dbName).FormatDSN()
}

// We're loading the driver anonymously, aliasing its
// package qualifier to _ so none of its exported names are
// visible to our code.
// Under the hook, the driver registers itself as being
// available to the database/sql package.
//
// ## Accessing the Database
// To create a sql.DB, use sql.Open(). This returns a *sql.DB.
// sql.Open() does not establish any connections to the database.
// THe first actual connection to the underlying datastore
// will be established lazily, when it's needed for the
// first time.
func OpenMyDB(c Conn, dbName string) (*sql.DB, error) {

	db, err := sql.Open("mysql", BuildDSN(c, dbName).FormatDSN())
	if err != nil {
		return nil, err
	}

	// If you want to check right away that the database is
	// available and accessible, use db.Ping().
	if err := db.Ping(); err != nil {
		return nil, err
	}

	// See https://github.com/go-sql-driver/mysql/issues/674
	db.SetConnMaxIdleTime(time.Second)

	// defer db.Close()
	// Call db.Close() if the sql.DB should not have a lifetime
	// beyond the scope of the function.
	// The sql.DB object is designed to be long-lived.
	// Don't Open() and Close() database frequently.
	// Create one sql.DB object for each distinct datastore
	// you need to access, and keep it until the program
	// is done accessing that datastore.
	// Pass it around as needed.
	return db, nil
}

func MustOpenMyDB(dbName string) *sql.DB {
	db, err := OpenMyDB(MustLoad().GetMySQLConn(), dbName)
	if err != nil {
		panic(err)
	}

	return db
}

func createDBStmt(name string) string {
	return fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %s
	CHARACTER SET utf8mb4 
	COLLATE utf8mb4_unicode_ci;`, name)
}

func CreateDB(db *sql.DB, dbName string) (string, error) {
	stmt := createDBStmt(dbName)
	_, err := db.Exec(stmt)
	if err != nil {
		return "", err
	}

	return stmt, nil
}
