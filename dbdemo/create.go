package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name     string
	Age      int
	Birthday time.Time
	gorm.Model
}
