package dao

import (
	"github.com/gogoclouds/go-gorm/model"
)

// email 没有组合 gorm.model 时， 没有软删除等

// 不想引入 gorm.Model, 这样启用软删除特性
// type User struct {
// 	ID int
//	Deleted gorm.DeletedAt
//	Name string
// }

// 删除一条记录时，需要指定条件，否则会触发批量删除
func (dao *UserDao) delete() {
	e := &model.Email{ID: 10}
	// delete from emails where id = 10;
	dao.db.Delete(e)

	// delete from emails where id = 10 and name = 'su.xu';
	dao.db.Where("name = ?", "su.xu").Delete(e)

	dao.db.Delete(e, 10)             // delete from emails where id = 10;
	dao.db.Delete(e, "10")           // delete from emails where id = 10;
	dao.db.Delete(e, []int{1, 2, 3}) // delete from emails id in (1, 2, 3);
}

func (dao *UserDao) deletedAt() {
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}
	// update users set daleted_at = "2022-11-08 22:50:02" where id = "xxxx-xxx-xx-xxx";
	dao.db.Delete(u)

	// select * from users where age = 20 and deleted_at IS NULL;
	dao.db.Where("age = 20").Find(u)
}

func (dao *UserDao) deleteUnscoped() {
	u := &model.User{ID: "xxxx-xxx-xx-xxx"}

	// 查询所有（包括软删除）
	// select * from users where age = 20;
	dao.db.Unscoped().Where("age = 20").Find(u)

	// 永久删除
	// delete from users where id = "xxxx-xxx-xx-xxx";
	dao.db.Unscoped().Delete(u)
}
