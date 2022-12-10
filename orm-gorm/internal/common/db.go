package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		CreateBatchSize: 1000, // 批量插入每次拆成 1k 条
		QueryFields:     true, // 会根据当前model的所有字段名称进行 select
	})
}
