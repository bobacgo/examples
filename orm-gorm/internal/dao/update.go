package dao

import (
	"log"

	"github.com/gogoclouds/orm-gorm/internal/model"
	"gorm.io/gorm"
)

func (dao *UserDao) Save() {
	var user *model.User
	user.Name = "lanjin.wei"
	user.Age = 100
	// 会保存所有的字段，即使是零值
	dao.db.Save(user)
}

// 单个字段更新
func (dao *UserDao) Update() {
	// update users set name = "xi.wei", updated_at = "2022-11-07 22:22:00" where active = true;
	dao.db.Model(&model.User{}).Where("active = ?", true).Update("name", "xi.wei")

	// update users set name = "yongli.chen", updated_at = "2022-11-07 22:22:00" where id = "xxxx";
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}
	dao.db.Model(u).Update("name", "yongli.chen")
}

// 多字段更新 -- 非零值
// 根据 `strut` 更新属性，只会更新非零值的字段
func (dao *UserDao) Updates() {
	// update users set name = "hui.wei", age = 18, updated_at = "2022-11-07 22:22:00" where id = "xxx";
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}
	dao.db.Model(u).Updates(model.User{Name: "hui.wei", Age: 18, Active: false})
}

// 多字段更新 -- 零值
func (dao *UserDao) UpdatesBySelectAndOmit() {
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}

	// select 所有字段 - 包含零值
	dao.db.Model(u).Select("Name", "Age").Updates(model.User{Name: "hongmei.wei", Age: 0})
	// select 除 role 外的所有字段
	dao.db.Model(u).Select("*").Omit("Role").Updates(model.User{Name: "cao.cao", Role: "admin", Age: 0})
}

func (dao *UserDao) UpdateByGlobal() {
	// 没有任何条件下，更新
	err := dao.db.Model(&model.User{}).Update("name", "bei.li").Error
	log.Println(err) // gorm.ErrMissingWhereClause

	// update users set name = 'yu.zhou' where 1 = 1
	dao.db.Model(&model.User{}).Where("1 = 1").Update("name", "yu.zhou")
}

// 使用sql表达更新
func (dao *UserDao) UpdateByExpr() {
	p := &model.Product{
		Model: gorm.Model{ID: 1},
	}

	// update product set price = price * 2 + 100, updated_at = "2022-11-07 22:22:00" where id = 1;
	dao.db.Model(p).Update("price", gorm.Expr("price * ? + ?", 2, 100))

	// update product set quantity = quantity - 1 where  id = 1 and quantity > 1;
	dao.db.Model(p).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
}

// 根据子查询进行更新
func (dao *UserDao) UpdateSubSelect() {
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}

	// update users set company_name = (select name from companies where companies.id = users.company_id);
	dao.db.Model(u).Update("company_name", dao.db.Model(&model.Company{})).Select("name").Where("companies.id = users.company_id")
	dao.db.Table("users as u").Where("name = ?", "yun.zhao").
		Update("company_name", dao.db.Table("companies as c").Select("name").Where("c.id = u.company_id"))
}
