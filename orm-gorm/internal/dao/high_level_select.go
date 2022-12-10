package dao

import (
	"github.com/gogoclouds/orm-gorm/internal/model"
	"gorm.io/gorm/clause"
)

// 高级查询

type APIUser struct {
	ID   string
	Name string
}

func (dao *UserDao) Find_fields() {
	// select id, name from users limit 10;
	dao.db.Model(&model.User{}).Limit(10).Find(&APIUser{})
}

// locking
func (dao *UserDao) Find_lock() {
	var users []*model.User

	// select * from users for update;
	dao.db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(users)

	// select * from users for share of users;
	dao.db.Clauses(clause.Locking{Strength: "SHARE", Table: clause.Table{Name: clause.CurrentTable}}).Find(users)

	// select * from users for update nowait;
	dao.db.Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).Find(users)
}

// 组合查询
func (dao *UserDao) Find_whereGroup() {
	d := dao.db

	// select * from pizza where (pizza = 'pepperoni' and (size = 'small' or size = 'medium')) or (pizza = 'hawaiian' and size = 'xlarge');
	d.Where(d.Where("pizza = ?", "pepperoni").Where(d.Where("size = ?", "small").Or("size = ?", "medium"))).
		Or(d.Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge")).
		Find(&model.Pizza{})
}
