package sq

import (
	"database/sql"
	"goplay/web/chrono"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// gorm.Model documentations are scattered around multiple places:
// * https://gorm.io/docs/models.html#gorm-Model
// * https://gorm.io/docs/models.html#Creating-x2F-Updating-Time-x2F-Unix-Milli-x2F-Nano-Seconds-Tracking
// * https://gorm.io/docs/conventions.html#Timestamp-Tracking
// * https://gorm.io/docs/gorm_config.html#NowFunc
type Product struct {
	gorm.Model // Not usable when you want to output to JSON.
	Code       string
	Price      uint
}

// Define Models
// By default, GORM uses:
// * `ID` field as primary key,
// * snake form of plural struct name as table name,
// * CreatedAt, UpdatedAt to track creation, update time.

type User struct {
	ID           string         `json:"id" gorm:"type:VARCHAR(36);primakeyKey"`
	Name         string         `json:"name" gorm:"type:VARCHAR(64)"`
	Email        string         `json:"email" gorm:"type:VARCHAR(64);uniqueIndex"`
	Age          uint           `json:"age"`
	Birthday     chrono.Date    `json:"birthday" gorm:"type:DATE"`
	MemberNumber sql.NullString `json:"memberNumber" gorm:"type:VARCHAR(36)"`
	ActivateAt   sql.NullTime   `json:"activateAt" gorm:"type:DATETIME(0)"`
	CreatedAt    time.Time      `json:"createdAt" gorm:"type:DATETIME(0);autoCreateTime"`
	UpdatedAt    time.Time      `json:"updatedAt" gorm:"type:DATETIME(0);autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"type:DATETIME(0);index"`
}

func NewUser() User {
	return User{
		ID:       uuid.New().String(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Age:      uint(gofakeit.Number(1, 130)),
		Birthday: chrono.DateNow(),
	}
}

// Field tags
// * `gorm:"column:cus_column_name"`
// * `gorm:"type:int"`
// * `gorm:"serializer:json/gob/unixtime"`
// * `gorm:"size:256"`
// * `gorm:"primaryKey"`
// * `gorm:"unique"`
// * `gorm:"defualt:xxx"`
// * `gorm:"precision"`
// * `gorm:"scale"`
// * `gorm:"notNull"`
// * `gorm:"autoIncrement"`
// * `gorm:"autoIncrementIncrement"`
// * `gorm:"embedded"`
// * `gorm:"embeddedPrefix"`
// * `gorm:"autoCreateTime"`
// * `gorm:"autoUpdateTime"`
// * `gorm:"index"`
// * `gorm:"uniqueIndex"`
// * `gorm:"check:age>13"`
// * `gorm:"<-:create,<-:update,<-:false"`
// * `gorm:"->:false"`
// * `gorm:"-,-:migration,-:all"`
// * `gorm:"comment"`

// Association Tags
// * `gorm:"foreignKey"` The column name of the current model
//    used as a foreign key in the join table
// * `gorm:"references"` The column name of the reference
//    table that the foreign key of the join table maps to.
// * `gorm:"polypmorphic"`
// * `gorm:"polymorphicValue"`
// * `gorm:"many2many"`
// * `gorm:"joinForeignKey"`
// * `gorm:"joinReferences"`
// * `gorm:"constraint"`

// CREATE TABLE credit_cards (
//
//		id          bigint unsigned AUTO_INCREMENT,
//		created_at  datetime(3) NULL,
//	 updated_at  datetime(3) NULL,
//	 deleted_at  datetime(3) NULL,
//	 number      longtext,
//	 customer_id bigint unsigned
//	 PRIMAEY KEY (id)
//	 INDEx       idx_credeit_cards_deleted_at (deleted_at)
//	 CONSTRAINT  fk_credit_cards_customer FOREIGN KEY (customer_id) REFERENCES customers (id)
//
// );
type CreditCard struct {
	gorm.Model
	Number     string
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID"`
}

// The example in the tutorial does not work.
type Customer struct {
	gorm.Model
	Name string
}

type Language struct {
	Code string
	Name string
}
