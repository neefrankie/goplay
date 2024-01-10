package orm

import (
	"goplay/web/model"
	"testing"
)

// When deleting a record, the deleted value should have
// primary key or a Batch Delete will be triggered.
func TestDeleteRecord(t *testing.T) {
	db := getMyDB()

	user := model.User{
		// ID: 111,
	}

	db.Delete(&user)
	// DELETE from users WHERE id = 111;

	db.Where("name = ?", "jinzhu").Delete(&user)
	// DELETE FROm users
	// WHERE id=111 AND name="jinzhu"
}

func TestDeleteWithPrimaryKey(t *testing.T) {
	db := getMyDB()

	db.Delete(&model.User{}, 10)
	// DELETE FROM users
	// WHERE id=10

	db.Delete(&model.User{}, "10")
	// DELETE FROM users
	// WHERE id=10
}

func TestBatchDelete(t *testing.T) {
	db := getMyDB()

	db.Where("name LIKE ?", "%jinzhu%").Delete(&model.User{})
	// DELETE FROM emails
	// WHERE email LIKE "%jinzhu%"

	db.Delete(&model.User{}, "email LIKE ?", "%jinzhu%")
	// DELETE FROM emails
	// WHERE email LIKE "%jinzhu%"

	// To efficiently delete large number of records:
	var users = []model.User{
		// {ID: 1},
		// {ID: 2},
		// {ID: 3},
	}
	db.Delete(&users)
	// DELETE FROM users WHERE id IN (1,2,3);

	db.Delete(&users, "name LIKE ?", "%jinzhu%")
	// DELETE FROM users WHERE name LIKE "%jinzhu%" AND id IN (1, 2, 3)
}

// Soft delete only works with gorm.DeletedAt field.
func TestSoftDelete(t *testing.T) {
	db := getMyDB()

	user := model.User{
		// ID: 111,
	}

	db.Delete(&user)
	// UPDATE users
	// SET deleted_at='xxxx'
	// WHERE id=111
}
