package dao

import (
	"github.com/gogoclouds/go-gorm/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserDao struct {
	db *gorm.DB
}

// db.Create(&user)
// db.Select("Name", "Age", "CreatedAt").Create(&user)
// db.Omit("Name", "Age", "CreateAt").Create(&user)
// db.CreateInBatches(users, 100)

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) Create() {
	user := model.User{Name: "Lanjin.wei", Age: 18, Birthday: time.Now()}
	result := dao.db.Create(&user)
	log.Println("id: ", user.ID)               // 返回插入数据的主键
	log.Println("err: ", result.Error)         // 返回 error
	log.Println("rows: ", result.RowsAffected) // 返回插入记录的条数
}

// 指定字段创建记录
func (dao *UserDao) CreateByFields() {
	user := model.User{Name: "Lanjin.wei", Age: 18, Birthday: time.Now()}
	// insert into users (name, age, create_at) values ("lanjin.wei", 18, "2020-07-04 11:05:21.775")
	dao.db.Select("Name", "Age", "CreatedAt").Create(&user)
	// 取反
	// insert into users (birthday, updated_at) values ("2020-07-04 11:05:21.775", "2020-07-04 11:05:21.775")
	dao.db.Omit("Name", "Age", "CreateAt").Create(&user)
}

// 批量插入
func (dao *UserDao) CreateInBatches() {
	// 数据量较少的情况
	users := []model.User{{Name: "yongli.chen"}, {Name: "xi.wei"}, {Name: "lanjin.wei"}}
	dao.db.Create(&users)

	// 插入时可以指定每批的数量 （可以在连接数据库的时候指定）
	users = []model.User{{Name: "yongli.cheng"}, {Name: "lanjin.wei"}}
	dao.db.CreateInBatches(users, 100)
}
