package db

import (
	"goplay/web/config"
	"testing"
)

func TestCreator_CreateDB(t *testing.T) {
	names := []string{
		"gormdb",
		"entdemo",
	}

	c := MustNewCreator(config.MustLoad().GetMySQLConn())

	for _, name := range names {
		stmt, err := c.CreateDB(name)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s\n", stmt)
	}
}
