package dao_test

import (
	"context"
	"log"
	"testing"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ud := dao.NewUserDao(g.DB.WithContext(ctx))
	ud.Create()
	ud.CreateByFields()
	ud.CreateInBatches()
}
