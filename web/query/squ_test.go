package query

import (
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestSqSelect(t *testing.T) {
	users := squirrel.Select("*").From("users").Join("emails USING (email_id)")

	active := users.Where(squirrel.Eq{"deleted_at": nil})

	sql, args, err := active.ToSql()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("SQL: %s\n", sql)
	t.Logf("Args: %v\n", args...)
}

func TestSqInsert(t *testing.T) {
	sql, args, err := squirrel.
		Insert("users").Columns("name", "age").
		Values("moe", 13).Values("larry", squirrel.Expr("? + 5", 12)).
		ToSql()

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s\n", sql)
	t.Logf("%v\n", args)
}
