package web

import (
	"goplay/web/config"
	"testing"
)

func TestMySQLCreateDB(t *testing.T) {
	names := []string{
		"gormdb",
		"entdemo",
	}

	db := config.MustOpenMyDB("")

	for _, name := range names {
		stmt, err := CreateTable(db, name)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s\n", stmt)
	}
}
