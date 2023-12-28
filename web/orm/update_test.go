package orm

import (
	"testing"

	"gorm.io/gorm"
)

func TestSaveAllFields(t *testing.T) {
	db := getMyDB()

	var user User
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	// UPDATE users
	// SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// Save combines Created and Update.
	// It performs Save if the the field for primary key is not set.
	// Don't use Save with Model together.
}

func TestUpdateSingleCol(t *testing.T) {
	db := getMyDB()

	// When updating a single column with Update, conditions
	// are required, otherwise ErrMissingWhereClause will be raised.
	// When using the Model method and its value has a primary
	// value, the primary key will be used to build the condition:
	db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users
	// SET name = 'hello', updated_at='xxxxx'
	// WHERE active=true

	user := User{
		// ID: 111,
	}
	db.Model(&user).Update("name", "hello")
	// UPATE users
	// SET name='hello', updated_at='xxxx'
	// WHERE id=111

	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users
	// SET name='hello', updated_at='xxx'
	// WHERE id=111 AND active=true
}

// DB.UPdates supports updating with struct or map[string]interface{}
// When updating with struct, it will only update non-zero fields by default.
// Use map or Select to specify zero fields to update.
func TestUpdateMultiCols(t *testing.T) {
	db := getMyDB()

	user := User{
		// ID: 111,
	}
	db.Model(&user).Updates(User{
		Name: "hello",
		Age:  18,
	})
	// UPDATE users
	// SET name='hello', age=18, updated_at='xxxx'
	// WHERE id=111

	db.Model(&user).Updates(map[string]interface{}{
		"name":   "hello",
		"age":    18,
		"active": false,
	})
	// UPDATE users
	// SET name='hello', age=18, active=false, updated_at='xxxx'
	// WHERE id=111
}

// If you want to update selected fields or ignore some fields,
// use Select, Omit
func TestUpdateSelectedFields(t *testing.T) {
	db := getMyDB()

	user := User{
		// ID: 111,
	}

	db.Model(&user).Select("name").Updates(map[string]interface{}{
		"name":   "hello",
		"age":    18,
		"active": false,
	})
	// UPDATE users
	// SET name='hello'
	// WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{
		"name":   "hello",
		"age":    18,
		"active": false,
	})
	// UPDATE users
	// SET age=18, active=false, updated_at='xxxx'
	// WHERE id=111;

	db.Model(&user).Select("Name", "Age").Updates(User{
		Name: "new_name",
		Age:  0,
	})
	// UPDATE users
	// SET name='new_name', age=0
	// WHERE id=111;

	// Select all fields, including zero fields.
	db.Model(&user).Select("*").Updates(User{
		Name: "jinzhu",
		Age:  0,
	})

	// Select all fields but omit Role
	db.Model(&user).Select("*").Omit("Role").Updates(User{
		Name: "jinzhu",
		Age:  0,
	})
}

// If you haven't specified a primary key with Model,
// a batch update will be performed.
func TestBatchUpdates(t *testing.T) {
	db := getMyDB()

	db.Model(User{}).Where("role = ?", "admin").Updates(User{
		Name: "hello",
		Age:  18,
	})
	// UPDATE users
	// SET name='hello', age=18
	// WHERE role='admin';

	db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{
		"name": "hello",
		"age":  18,
	})
}

func TestUpdateResult(t *testing.T) {
	db := getMyDB()

	result := db.Model(User{}).Where("role=?", "admin").Updates(User{
		Name: "hello",
		Age:  18,
	})

	t.Logf("RowsAffected %d", result.RowsAffected)
}

func TestUpdateWithSQLExpr(t *testing.T) {
	db := getMyDB()

	product := Product{
		Model: gorm.Model{
			ID: 3,
		},
	}

	db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "products"
	// SET "price" = price * 2 + 100, updated_at='xxxx'
	// WHERE id=3
}
