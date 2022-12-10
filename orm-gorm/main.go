package main

import (
	"log"

	"github.com/gogoclouds/orm-gorm/internal/common"
	"github.com/gogoclouds/orm-gorm/internal/common/g"
	"github.com/gogoclouds/orm-gorm/internal/model"
)

func main() {
	var err error
	// 连接数据库
	if g.DB, err = common.ConnectDB(); err != nil {
		log.Fatal("gorm.Open(): ", err)
	}
	// 创建表
	model.CreateTable(g.DB)
}
