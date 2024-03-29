package models

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	DB          *gorm.DB
	GormAdapter *gormadapter.Adapter
)

type BaseMigrate struct {
	ID        uint      `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseModel struct {
	ID        uint      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func DBConnectInit(dbStr string) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	log.Fatalf("connect db failed, err: %v \n", err)
	//	return
	//}
	//db.SetMaxOpenConns(200)
	//db.SetMaxIdleConns(100)
	//DB = db
	db, err := gorm.Open("mysql", fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dbStr))
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
	//GormAdapter = gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test", true)
	//GormAdapterPool = gormAdapter.NewAdapter("mysql", "root:123456@(127.0.0.1:3306)/test", true)
	GormAdapter = gormadapter.NewAdapterByDB(DB)
	return
}
