package dao

import "github.com/gogoclouds/go-gorm/model"

// 子查询

type AvgAgeResult struct {
	Name   string
	AvgAge uint16
}

func (dao *UserDao) Find_subselect() {
	var order []*model.Order

	// select * from orders where amount > (select avg(amount) from orders);
	dao.db.Where("amount > (?)", dao.db.Table("orders").Select("avg(amount)")).Find(order)

	var avgAgeResults []*AvgAgeResult

	/*
		select name, avg(age) as avgage from users group by name having avg(age) > (
			select avg(age) from users where name like "%.zhang"
		);
	*/
	subq := dao.db.Select("avg(age)").Where("name like ?", "%.zhang").Table("users") // sub query
	dao.db.Select("avg(age) as avgage").Group("name").Having("avg(age) > (?)", subq).Find(avgAgeResults)

	var user *model.User
	// select * from (select name, age from users) as u where age = 18
	dao.db.Table("(?) as u", dao.db.Model(&model.User{}).Select("name", "age")).Where("age = ?", 18).Find(user)
}
