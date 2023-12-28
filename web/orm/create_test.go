package orm

import (
	"goplay/web/chrono"
	"testing"
)

func TestCreateOne(t *testing.T) {
	db := getMyDB()
	db.AutoMigrate(&User{})

	user := User{
		Name:     "Jinzhu",
		Age:      18,
		Birthday: chrono.DateNow(),
	}

	result := db.Create(&user)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Logf("Created user %s", user.ID)
	t.Logf("Rows affected: %d", result.RowsAffected)
}

func TestCreateMulti(t *testing.T) {
	db := getMyDB()

	users := []*User{
		{
			Name:     "Jinzhu",
			Age:      18,
			Birthday: chrono.DateNow(),
		},
		{
			Name:     "Jackson",
			Age:      19,
			Birthday: chrono.DateNow(),
		},
	}

	result := db.Create(users)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Logf("Inserted %d rows\n", result.RowsAffected)
}

// When creating from map, primary key values won't be back filled.
func TestCreateMap(t *testing.T) {
	db := getMyDB()

	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu",
		"Age":  18,
	})

	db.Model(&User{}).Create([]map[string]interface{}{
		{
			"Name": "jinzhu_1",
			"Age":  18,
		},
		{
			"Name": "jinzhu_2",
			"Age":  20,
		},
	})
}
