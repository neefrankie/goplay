package orm

import "testing"

// The Count method is used to retrieve the number of records
// that match a given query.
func TestCount(t *testing.T) {
	db := getMyDB()

	var count int64
	db.Model(&User{}).
		Where("name = ?", "jinzhu").
		Or("name = ?", "jinzhu 2").
		Count(&count)
}
