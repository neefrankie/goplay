package orm

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestRawSQL(t *testing.T) {
	var u1 User

	db := getMyDB()
	db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&u1)

	t.Logf("%v", u1)
}

func TestRawExec(t *testing.T) {
	db := getMyDB()

	result := db.Exec("DROP TABLE users")
	if result.Error != nil {
		t.Error(result.Error)
	}
}

func TestExecUpdate(t *testing.T) {
	db := getMyDB()

	result := db.Exec("UPDATE orders SET shipped_at = ? WHERE IN ?", time.Now(), []int64{1, 2, 3})
	if result.Error != nil {
		t.Error(result.Error)
	}
}

func TestExecWithSQLExpr(t *testing.T) {
	db := getMyDB()

	db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")
}
