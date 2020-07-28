package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"job_spider/model"
)

var DB *gorm.DB

func DBInit()  {
	db,err := gorm.Open("mysql","root:123456@(localhost)/job?charset=utf8mb4&parseTime=True&loc=Local")
	if(err != nil){
		fmt.Println("打开数据库失败"+err.Error())
		panic("打开数据库失败"+err.Error())
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(5)
	db.DB().SetMaxOpenConns(20)
	db.AutoMigrate(
		&model.Job{},
	)
	DB = db
}
