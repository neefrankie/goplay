package config

import "testing"

func TestCreateDB(t *testing.T) {
	names := []string{
		"gormdb",
		"entdemo",
	}

	db := MustOpenMyDB("")

	for _, name := range names {
		stmt, err := CreateTable(db, name)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%s\n", stmt)
	}
}
