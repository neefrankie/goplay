package query

import (
	"testing"

	"github.com/doug-martin/goqu/v9"
)

func TestQuSelect(t *testing.T) {
	sql, _, _ := goqu.From("test").ToSQL()

	t.Logf("%s\n", sql)
}

func TestQuInsert(t *testing.T) {
	ds := goqu.Insert("user").
		Cols("first_name", "last_name").
		Vals(
			goqu.Vals{"Greg", "Farley"},
			goqu.Vals{"Jimmy", "Stewart"},
			goqu.Vals{"Jeff", "Jeffers"},
		)

	insertSQL, args, _ := ds.ToSQL()
	t.Logf("%s\n", insertSQL)
	t.Logf("%v\n", args)
}

func TestQuInsertRecord(t *testing.T) {
	ds := goqu.Insert("user").Rows(
		goqu.Record{"first_name": "Greg", "last_name": "Farley"},
		goqu.Record{"first_nname": "Jimmy", "last_name": "Stewart"},
		goqu.Record{"first_name": "Jeff", "last_name": "Jeffers"},
	)

	insertSQL, args, _ := ds.ToSQL()
	t.Logf("%s\n", insertSQL)
	t.Logf("%v\n", args)
}

func TestQuInsertStruct(t *testing.T) {
	type User struct {
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
	}

	ds := goqu.Insert("user").Rows(
		User{FirstName: "Greg", LastName: "Farley"},
		User{FirstName: "Jimmy", LastName: "Stewart"},
		User{FirstName: "Jeff", LastName: "Jeffers"},
	)

	insertSQL, args, _ := ds.ToSQL()
	t.Logf("%s\n", insertSQL)
	t.Logf("%v\n", args)
}
