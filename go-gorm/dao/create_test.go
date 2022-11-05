package dao_test

import (
	"log"
	"testing"

	"github.com/gogoclouds/go-gorm/common"
	"github.com/gogoclouds/go-gorm/common/g"
	"github.com/gogoclouds/go-gorm/dao"
)

func init() {
	var err error
	g.DB, err = common.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	// model.CreateTable(g.DB)
}

func TestCreate(t *testing.T) {
	ud := dao.NewUserDao(g.DB)
	ud.Create()
	ud.CreateByFields()
	ud.CreateInBatches()
}
