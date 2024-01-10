package orm

import (
	"testing"
)

func TestCreateOne(t *testing.T) {
	db := getMyDB()
	db.AutoMigrate(&User{})

	user := NewUser()

	result := db.Create(&user)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Logf("Created user %d", user.ID)
	t.Logf("Rows affected: %d", result.RowsAffected)
}

func TestCreateMulti(t *testing.T) {
	db := getMyDB()

	users := []*User{
		NewUserP(),
		NewUserP(),
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

	db.Model(&User{}).Create(NewUserM())

	db.Model(&User{}).Create([]map[string]interface{}{
		NewUserM(),
		NewUserM(),
	})
}
