package main

import "testing"

func TestCreateRecord(t *testing.T) {
	myDB := mustConnect()

	myDB.AutoMigrate(&User{})

}
