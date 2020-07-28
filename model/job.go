package model

import "time"

type Job struct {
	ID 			int64 			`gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name 		string			`gorm:"NOT NULL;content:'职位名'"`
	Company 	string			`gorm:"NOT NULL;content:'公司名'"`
	Address 	string			`gorm:"NOT NULL;content:'公司地址'"`
	MinSalary 	int64			`gorm:"NOT NULL;content:'最小薪资'"`
	MaxSalary 	int64			`gorm:"NOT NULL;content:'最大薪资'"`
	Class 		string			`gorm:"NOT NULL;content:'分类'"`
	Status		int				`gorm:"NOT NULL;default:1;content:'状态0下架1上架'"`
	PublishAt	time.Time
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}
