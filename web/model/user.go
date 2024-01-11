package model

import (
	"database/sql"
	"goplay/web/chrono"
	"goplay/web/rbac"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string         `gorm:"type:VARCHAR(64)"`
	Email         string         `gorm:"type:VARCHAR(64);uniqueIndex" validate:"required,email"`
	Age           uint8          `validate:"gte=0,lte=130"`
	Role          rbac.Role      `gorm:"type:VARCHAR(16)"`
	Birthday      chrono.Date    `gorm:"type:DATE"`
	MemberNumber  sql.NullString `gorm:"type:VARCHAR(36)"`
	ActivateAt    sql.NullTime   `gorm:"type:DATETIME(0)"`
	FirstName     string         `validate:"required"`
	LastName      string         `validate:"required"`
	Gender        string         `validate:"oneof=male female prefer_not_to"`
	FavoriteColor string         `validate:"iscolor"`
	Addresses     []*Address     `validate:"required,dive,required"`
}

type Users []User

func (u Users) Exists(id int) bool {
	for _, user := range u {
		if user.ID == uint(id) {
			return true
		}
	}

	return false
}

func (u Users) FindByName(name string) (User, error) {
	for _, user := range u {
		if user.Name == name {
			return user, nil
		}
	}

	return User{}, sql.ErrNoRows
}

func CreateUsers() Users {
	users := Users{}

	users = append(users, User{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "Admin",
		Role: "admin",
	})
	users = append(users, User{
		Model: gorm.Model{
			ID: 2,
		},
		Name: "Sabine",
		Role: "member",
	})
	users = append(users, User{
		Model: gorm.Model{
			ID: 3,
		},
		Name: "Sepp",
		Role: "member",
	})

	return users
}

type Address struct {
	gorm.Model
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())
