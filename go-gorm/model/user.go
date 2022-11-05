package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primarykey"`
	Name     string
	Role     string
	Age      uint8
	Birthday time.Time
}

// hooks
// BeforeSave BeforeCreate AfterSave AfterCreate
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// add uuid
	u.ID = uuid.NewString()
	if !u.IsValid() {
		return errors.New("can't save invalid data")
	}
	return
}

func (u *User) IsValid() bool {
	// u.Role == "" check role
	// ...
	return true
}
