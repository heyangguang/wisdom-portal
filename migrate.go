package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"wisdom-portal/models"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect db failed, err: %v \n", err)
		return
	}
	defer db.Close()
	db.SingularTable(true)

	//初始化数据库
	db.AutoMigrate(&models.User{}, &models.UserGroup{})
}
