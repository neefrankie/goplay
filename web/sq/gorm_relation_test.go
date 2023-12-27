package sq

import "testing"

func TestForeignKey(t *testing.T) {
	db := getMyDB()

	db.AutoMigrate(&Customer{}, &CreditCard{})
}
