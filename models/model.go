package models

import (
	gormAdapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	DB          *gorm.DB
	GormAdapter *gormAdapter.Adapter
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DBConnectInit() {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	log.Fatalf("connect db failed, err: %v \n", err)
	//	return
	//}
	//db.SetMaxOpenConns(200)
	//db.SetMaxIdleConns(100)
	//DB = db
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("connect db failed, err: %v \n", err)
		return
	}
	//defer db.Close()
	DB = db
	DB.SingularTable(true)
	SqlxAdapterInit()
	return
}

func SqlxAdapterInit() {
	// 设置连接配置，这里不做连接动作
	//opts := &sqlxadapter.AdapterOptions{
	//DriverName:     "mysql",
	//DataSourceName: "root:123456@tcp(127.0.0.1:3306)/test",
	//TableName:      "SYS_USER_GROUP_ROLE",
	//DB:             DB,
	//}
	// 这里真正做了数据库的连接
	//SqlxAdapter = sqlxadapter.NewAdapterFromOptions(opts)
	//GormAdapter = gormAdapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test", true)
	GormAdapter = gormAdapter.NewAdapterByDB(DB)
	return
}
