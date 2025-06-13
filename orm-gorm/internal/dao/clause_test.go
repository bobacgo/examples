package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUpdate(t *testing.T) {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/exp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	CreateUpdate(db)
}
