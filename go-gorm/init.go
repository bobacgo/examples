package main

import (
	"log"

	"github.com/gogoclouds/go-gorm/common"
	"github.com/gogoclouds/go-gorm/common/g"
	"github.com/gogoclouds/go-gorm/model"
)

func init() {
	var err error
	// 连接数据库
	if g.DB, err = common.ConnectDB(); err != nil {
		log.Fatal("gorm.Open(): ", err)
	}
	// 创建表
	model.CreateTable(g.DB)
}
