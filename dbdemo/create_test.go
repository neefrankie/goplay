package main

import (
	"testing"
	"time"
)

func TestCreateRecord(t *testing.T) {
	myDB := mustConnect()

	myDB.AutoMigrate(&User{})

	user := User{
		Name:     "Jinzhu",
		Age:      18,
		Birthday: time.Now(),
	}

	result := myDB.Create(&user)

	if result.Error != nil {
		t.Fatal(result.Error)
	}
	t.Logf("ID %d, rows affected %d", user.ID, result.RowsAffected)

	users := []*User{
		{
			Name:     "Jinzhu",
			Age:      18,
			Birthday: time.Now(),
		},
		{
			Name:     "Jackson",
			Age:      19,
			Birthday: time.Now(),
		},
	}

	result = myDB.Create(users)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Logf("Rows affected %d", result.RowsAffected)
}
