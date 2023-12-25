package sq

// To use database/sql, you'll need the package itsel,
// as well as a driver for the specific database you want to use.
// You generally shouldn't use dirver package directly.
// Your code should only refer to types defined in database/sql.

import (
	"database/sql"
	"fmt"
	"goplay/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
func OpenDB(c config.Conn, dbName string) (*sql.DB, error) {

	db, err := sql.Open("mysql", buildDSN(c, dbName).FormatDSN())
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

func FetchUser(db *sql.DB) error {
	var id int
	var name string

	// Use db.Query to send the query to the database.
	rows, err := db.Query("SELECT id, name FROM users WHERE id = ?", 1)
	if err != nil {
		return err
	}

	defer rows.Close()
	// Iterate over the rows with rows.Next()
	for rows.Next() {
		// Read teh columns in each row into variables.
		// you can't get a row as a map.
		// You need to create variables of the correct type,
		// and pass pointer to them.
		// When you iterate over rows and scan them into
		// destination variables, Go performs data type
		// conversions for you.
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		log.Println(id, name)
	}
	// You should always check for an error at the end of
	// the for row.Next() loop.
	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

// You should always prepare queries to be used multiple times.
// Under the hood, db.Query actually prepares, executes,
// and closes a prepared statement.
func PrepareQueries(db *sql.DB) error {
	var id int
	var name string

	stmt, err := db.Prepare("SELECT id, name from users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(1)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

// If a query returns at most one row, you can use a shortcut.
func SingleRowQueries(db *sql.DB) error {
	var name string
	// Errors from the query are deferred until Scan() is called.
	// You can also cal QueryRow() on a prepared statement.
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}

// ## Statement that Modify Data
// Use Exec() to accomplish an INSERT, UPDATE, DELETE, or
// another statement that doesn't return rows.
// DO not use Query in such case:
// _, err := db.Query("DELETE FROM users")
// The Query() will return a sql.Rows, which reserves a
// database connection until the sql.Rows is closed.
// In this case, the connection will never be released again.
func InsertData(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec("Dolly")
	if err != nil {
		return err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return nil
}

// ## Transactions
// A transaction is essentially an object that reserves a
// connection to the datastore.
//
// Begin a transaxtion with a call to db.Begin(), and close it
// with a Commit() or Rollback() method on the resulting Tx.
//
// Under the covers, the Tx gets a connection from the pool,
// and reserves it for use only with that transaction.
// Prepared statements that are created in a transaction are
// bound execlusively to that transaction.
//
// You should not mingle the use of transaction-releated
// functions such as Begin() and Commit() with SQL statements
// such as BEGIN and COMMIT in SQL code.
//
// If you need to work with multiple statements that modify
// connection state, you need a Tx even if you don't want a
// transaction per se.

// ## Prepared Statements And Connections
// At the database level, a prepared statement is bound to a
// single database connection.
// Typical flow:
// 1. The client sends a SQL statement with placeholders to
//    the server for preparation;
// 2. The server responds with a statement ID;
// 3. The client executes the statement by sending its ID and
//    parameters.
//
// In Go, connections are not exposed directly to the user
// of the database/sql package.

// ## Avoiding Prepared Statements
//
// Go creates prepared statements for you under the covers.
// A simple db.QUery(sql, params1, param2) works by preparing
// the sql, then executing it with the parameters and finally
// closing the statement.
//
// If you don't want to use a prepared statement, use `fmt.
// Sprint()` or similar to assemble the SQL, and pass this
// as the only argument to db.Query() or db.QueryRow().
// And your driver needs to support plaintext query execution.

// ## Prepared Statement in Transactions
// Prepared statement that create in a Tx are bound
// exclusively to it. When you operate on a Tx object,
// your actions map directly to the one and only one
// connection underlying it.
// This means that prepared statement created inside a Tx
// can't be used separatedly from it.
// Prepared statements created on a DB can't be used within
// a transaction, because they will be bound to a
// different connection.
// To use a prepared statement prepared outside the
// transaction in a Tx, you can use Tx.Stmt(), which will
// create a new transaction-sepcific statement from the one
// prepared outside the transaction.

// ## Handling Errors
