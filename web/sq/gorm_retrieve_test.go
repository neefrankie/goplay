package sq

import (
	"encoding/json"
	"testing"
	"time"

	"gorm.io/gorm"
)

// To retrieve a singl row, use First, Take, Last.
// Return ErrRecordNotFound upon missing row.
// First and Last searchs by primary key.
// Dest must be struct.
// If a model has no primary key, records will be ordered by
// first field.
func TestGetFirstRow(t *testing.T) {
	db := getMyDB()
	// First record by primary key ascending.
	var u1 User
	result := db.First(&u1)
	// SELECT *
	// FROM users
	// ORDER BY users.id
	// LIMIT 1
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("First user: %v\n", u1)

	var u2 = map[string]interface{}{}
	// SELECT * FROM users ORDER BY users.id LIMIT 1
	result = db.Model(&User{}).First(&u2)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("User in map: %v", u2)

	// Not working
	// db.Table("users").First(&user)

	var u3 User
	// SELECT * FROM users WHERE id = 10
	result = db.First(&u3, 2)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("User 2: %v", u3)

	var u4 User
	// SELECT * FROM users WHERE id = 2
	// Only works for int id.
	// If the second param cannot be coverted to a string,
	// it is treated as a column in WHERE.
	// For UUID, don't use this approach.
	result = db.First(&u4, "2")
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("User 2: %v", u3)
}

func TestFirstWithoutPrimary(t *testing.T) {
	db := getMyDB()
	db.AutoMigrate(&Language{})

	db.Create(&Language{
		Code: "zh-cn",
		Name: "Chinese",
	})

	var l Language
	// SELECT * FROM languages ORDER BY languages.code LIMIT 1
	result := db.First(&l)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("%v\n", l)
}

func TestGetLastRow(t *testing.T) {
	db := getMyDB()

	// Last record by primary key descending.
	// If you reuse u1, it seems not retrieving the correct data.
	var u1 User
	result := db.Last(&u1)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("Last user: %v\n", u1)
}

func TestTakeAnyRow(t *testing.T) {
	db := getMyDB()

	// Get a row, no order.
	var u1 User
	result := db.Take(&u1)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("Any one of user: %v\n", u1)

	var u2 = map[string]interface{}{}
	result = db.Table("users").Take(&u2)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("Any user in amp: %v\n", u2)
}

func TestFind(t *testing.T) {
	db := getMyDB()

	var u123 []User
	// SELECT * FROM users WHERE id IN (1, 2, 3)
	result := db.Find(&u123, []int{1, 2, 3})
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("Found user 1 2 3: %v", u123)
}

// Select allows you to specify the fields you want to retrieve from db.
// Select all fields by default.
func TestSepcifyFields(t *testing.T) {
	db := getMyDB()

	var users []User
	db.Select("name", "age").Find(&users)
	// SELECT name, age
	// FROm users;

	db.Select([]string{"name", "age"}).Find(&users)
	// SELECT name, age
	// FROM users;
}

func TestStringCond(t *testing.T) {
	db := getMyDB()

	// Get first matched record
	var u1 User
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1
	result := db.Where("name = ?", "jinzhu").First(&u1)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("%v\n", u1)

	// Get all matched records
	var users []User
	// SELECT * FROM users WHERE name <> 'jinzhu'
	result = db.Where("name <> ?", "jinzhu").Find(&users)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("%v\n", users)

	// IN
	// SELECT * FROM users WHERE name IN ('jinzhu', 'jinzhu 2')
	db.Where("name IN ?", []string{"jinzhu", "jinhu 2"}).Find(&users)

	// LIKE
	// SELECT * FROM users WHERE name LIKE '%jin%'
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00'
	db.Where("updated_at > ?", "2000-01-01 00:00:00").Find(&users)

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", "2000-01-01 00:00:00", "2000-01-08 00:00:00").Find(&users)

	// If dest have primary key set, condition will not override it.
	// Instead they will be AND:
	// var user = User{ID: 10}
	// db.Where("id = ?", 20).First(&user)
	// SELECT * FROM users WHERE id = 10 AND id = 20 ORDER BY id ASC LIMIT 1
}

func TestStructMapCond(t *testing.T) {
	db := getMyDB()

	// Struct
	var u1 User
	db.Where(&User{Name: "jinzhu", Age: 28}).First(&u1)
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu" AND age = 20
	// ORDER BY id
	// LIMIT 1

	// Map
	var users []User
	db.Where(map[string]interface{}{
		"name": "jinzhu",
		"age":  20,
	}).Find(&users)
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu" AND age = 20

	// Slice of primary keys
	db.Where([]int64{20, 1, 22}).Find(&users)
	// SELECT *
	// FROM users
	// WHERE id IN (20, 21, 22)

	// When querying with struct, GORM will only query with non-zero fields.
	// Field's value with 0, '', false or other zero values wont't be used
	// to build query conditions.
	db.Where(&User{
		Name: "jinzhu",
		Age:  0,
	}).Find(&users)
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu"

	// To include zero values in the query conditions, use a map.
	db.Where(map[string]interface{}{
		"Name": "jinzhu",
		"Age":  0,
	}).Find(&users)
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu" AND age = 0;

	// When searching with struct, you can specify which particular values
	// from the struct to use in the query conditions by passing in the
	// relevant field name or the dbname:
	db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu" AND age = 0;

	db.Where(&User{
		Name: "jinzhu",
	}, "Age").Find(&users)
	// SELECT *
	// FROM users
	// WHERE age = 0;
}

// Query conditions can be inlined into methods like First and Find
// similar to Where.
func TestInelineCond(t *testing.T) {
	db := getMyDB()

	var user User
	db.First(&user, "id = ?", "string_primary_key")
	// SELECT *
	// FROM users
	// WHERE id = 'string_primary_key';

	var users []User
	db.Find(&users, "name = ?", "jinzhu")
	// SELECT *
	// FROM users
	// WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT *
	// FROM users
	// WHERE name <> "jinzhu" AND age > 20

	db.Find(&users, User{Age: 20})
	// SELECT *
	// FROM users
	// WHERE age = 20;

	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT *
	// FROM users
	// WHERE age = 20
}

func TestNotCond(t *testing.T) {
	db := getMyDB()

	var user User
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT *
	// FROM users
	// WHERE NOT name = "jinzhu"
	// ORDER BY id
	// LIMIT 1;

	var users []User
	db.Not(map[string]interface{}{
		"name": []string{"jinzhu", "jinzhu 2"},
	}).Find(&users)

	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT *
	// FROm users
	// WHERE name <> "jinzhu" AND age <> 18
	// ORDER BY id
	// LIMIT 1

	db.Not([]int64{1, 2, 3}).First(&user)
}

func TestOrCond(t *testing.T) {
	db := getMyDB()

	var users []User
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	// SELECT *
	// FROm users
	// WHERE role = 'admin' OR role = 'super_admin';

	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)

	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
}

func TestOrderBy(t *testing.T) {
	db := getMyDB()

	var users []User
	db.Order("age desc, name").Find(&users)
	// SELECT *
	// FROM users
	// ORDER BY age desc, name;

	db.Order("age desc").Order("name").Find(&users)
	// SELECT *
	// FROM users
	// ORDER BY age desc, name;
}

func TestLimitOffset(t *testing.T) {
	db := getMyDB()

	var users []User
	db.Limit(3).Find(&users)
	// SELECT *
	// FROM users
	// LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users).Limit(-1).Find(&users)

	db.Offset(3).Find(&users)
	// SELECT *
	// FROm users
	// OFFSET 3;

	db.Limit(10).Offset(5).Find(&users)
	// SELECT *
	// FROM users
	// OFFSET 5 LIMIT 10;
}

func TestDistinct(t *testing.T) {
	db := getMyDB()

	var results []User
	db.Distinct("name", "age").Order("name, age desc").Find(&results)
}

func TestJoins(t *testing.T) {
	db := getMyDB()

	type result struct {
		Name  string
		Email string
	}

	db.Model(&User{}).
		Select("user.name, emails.email").
		Joins("left join emails on emails.user_id = users.id").
		Scan(&result{})
}

func TestGormModelJSON(t *testing.T) {
	m := gorm.Model{
		ID:        1,
		CreatedAt: time.Now().Truncate(time.Second),
		UpdatedAt: time.Now().Truncate(time.Second),
	}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", b)
}
