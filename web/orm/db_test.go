package orm

import "testing"

func TestMigrate(t *testing.T) {
	db := getMyDB()

	db.AutoMigrate(&User{})
}
