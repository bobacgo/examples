package dao_test

import (
	"log"
	"testing"

	"github.com/gogoclouds/orm-gorm/internal/common"
	"github.com/gogoclouds/orm-gorm/internal/common/g"
	"github.com/gogoclouds/orm-gorm/internal/dao"
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
