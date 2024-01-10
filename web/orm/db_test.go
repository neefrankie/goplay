package orm

import (
	"goplay/web/model"
	"testing"
)

func TestMigrate(t *testing.T) {
	db := getMyDB()

	db.AutoMigrate(&model.User{})
}
