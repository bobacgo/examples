package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Record struct {
	RecDay string `gorm:"rec_day"`
	UserId int    `gorm:"user_id"`
	Total  int    `gorm:"total"`
}

func (r *Record) TableName() string {
	return "records"
}

func CreateUpdate(db *gorm.DB) {
	rd := Record{
		RecDay: "2023-01-03",
		UserId: 1,
		Total:  232,
	}

	// [3.509ms] [rows:0] INSERT INTO `records` (`rec_day`,`user_id`,`total`) VALUES ('2023-01-03',1,232) ON DUPLICATE KEY UPDATE `total`=VALUES(`total`)

	// tx := db.Debug().Save(&rd)
	// if tx.Error != nil {
	// 	panic(tx.Error)
	// }
	tx := db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "rec_day"}, {Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"total"}),
	}).Create(&rd)
	if err := tx.Error; err != nil {
		panic(err)
	}
}
