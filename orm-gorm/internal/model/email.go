package model

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	ID uint
}
