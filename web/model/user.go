package model

import (
	"database/sql"
	"goplay/web/chrono"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string         `gorm:"type:VARCHAR(64)"`
	Email         string         `gorm:"type:VARCHAR(64);uniqueIndex" validate:"required,email"`
	Age           uint8          `validate:"gte=0,lte=130"`
	Birthday      chrono.Date    `gorm:"type:DATE"`
	MemberNumber  sql.NullString `gorm:"type:VARCHAR(36)"`
	ActivateAt    sql.NullTime   `gorm:"type:DATETIME(0)"`
	FirstName     string         `validate:"required"`
	LastName      string         `validate:"required"`
	Gender        string         `validate:"oneof=male female prefer_not_to"`
	FavoriteColor string         `validate:"iscolor"`
	Addresses     []*Address     `validate:"required,dive,required"`
}

type Address struct {
	gorm.Model
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())
