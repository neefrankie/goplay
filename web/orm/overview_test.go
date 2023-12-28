package orm

import (
	"goplay/web/config"
	"testing"

	"gorm.io/gorm"
)

func getMyDB() *gorm.DB {
	return MustNewMySQL(config.MustLoad().GetMySQLConn())
}

func TestOverfiew(t *testing.T) {
	db := getMyDB()

	// Migrate schema
	db.AutoMigrate(&Product{})

	p1 := Product{
		Code:  "D42",
		Price: 100,
	}

	// Read
	// TODO:
	// * Where is error?
	// * Where is ResultSet?
	// * How to see the SQL emitted?
	db.Create(&p1)
	t.Logf("Inserted: %v", p1)

	var p2 Product
	db.First(&p2, p1.ID)
	t.Logf("Retrieved: %v", p2)

	db.First(&p2, "code = ?", "D42")
	t.Logf("Retrieved: %v", p2)

	// Update - change price to 200
	// TODO: Does the "Price" refer to struct field?
	// Or column name?
	db.Model(&p1).Update("Price", 200)

	// Update - multiple fields
	db.Model(&p1).Updates(Product{
		Price: 200,
		Code:  "F42",
	})
	db.Model(&p1).Updates(map[string]interface{}{
		"Price": 200,
		"Code":  "F42",
	})

	// Delete
	db.Delete(&p1, p1.ID)
}
