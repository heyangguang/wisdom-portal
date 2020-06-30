package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"wisdom-portal/models/migrate"
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
	db.Debug().AutoMigrate(&migrate.UserMigrate{}, &migrate.UserGroupMigrate{}, &migrate.RuleMigrate{}, &migrate.ObjActMigrate{},
		&migrate.CasbinRuleMigrate{}, &migrate.MonitorElasticSearchMigrate{}, &migrate.MonitorMySQLMigrate{}, &migrate.AlertMigrate{},
		&migrate.MonitorIntermediateMigrate{}, &migrate.MonitorKafkaMigrate{}, &migrate.MonitorKubernetesMigrate{},
		&migrate.MonitorAccessLogMigrate{})

	//objActSlice := make([]map[string]string, 0)
	//objActSlice = append(objActSlice, map[string]string {"objName": "/api/v1/perm", "actName": "GET", "tag": "perm"})
	//objActSlice = append(objActSlice, map[string]string {"objName": "/api/v1/perm", "actName": "POST", "tag": "perm"})
	//objActSlice = append(objActSlice, map[string]string {"objName": "/api/v1/user", "actName": "GET", "tag": "user"})
	//objActSlice = append(objActSlice, map[string]string {"objName": "/api/v1/*", "actName": "*", "tag": "admin"})
	//for _, value := range objActSlice {
	//	var objAct models.ObjAct
	//	objAct.ObjName = value["objName"]
	//	objAct.ActName = value["actName"]
	//	objAct.Tag = value["tag"]
	//	db.Debug().Create(&objAct)
	//}
}
