package model

import "gorm.io/gorm"

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
}
