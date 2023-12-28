package orm

import "testing"

func TestForeignKey(t *testing.T) {
	db := getMyDB()

	db.AutoMigrate(&Customer{}, &CreditCard{})
}
