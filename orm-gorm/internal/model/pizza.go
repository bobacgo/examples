package model

import "gorm.io/gorm"

type Pizza struct {
	gorm.Model
	Pizza string
	Size  string
}
