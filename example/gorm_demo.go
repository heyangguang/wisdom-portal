package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type UserLanguages struct {
//	ID string
//	ID string
//}

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
	gorm.Model
	Name string
}

func main() {
	db, _ := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Language{})

	//db.Create(&User{Name: "heyang"})
	//db.Create(&User{Name: "罗灿"})
	//db.Create(&User{Name: "何阳"})
	//db.Create(&User{Name: "郭锐"})
	//db.Create(&Language{Name: "中文"})
	//db.Create(&Language{Name: "英文"})
	//db.Create(&Language{Name: "俄语"})
	//var users []User
	//language := Language{}

	//db.Where("where id = ?", 1).First(&language)
	//var user User
	//var language []Language
	//db.Debug().Select("id").Where("name = ?", "何阳").First(&user)
	//db.Debug().Select("id").Find(&language)
	//fmt.Println(language)
	//for _, value := range *language {
	//	fmt.Println(value.ID)
	//}
	//fmt.Println(user.ID)
	//u1 := &User{
	//	Name: "何阳",
	//	Languages: []Language{
	//		{Name:"中文"},
	//		{Name:"英文"},
	//	},
	//}
	//db.Debug().Create(u1)

	// 2 -> 1 2 3
	// 多对多删除
	//db.Model(&User{Model: gorm.Model{ID: 1}}).Association("Languages").Clear()
	// 多对多插入
	//db.Model(&user).Association("Languages").Append(language)
	//err := db.Model(&User{Model: gorm.Model{ID: 1}}).Association("Languages").Append(&[]Language{
	//	{Model: gorm.Model{ID: 1}},
	//	{Model: gorm.Model{ID: 2}},
	//	{Model: gorm.Model{ID: 3}},
	//	{Model: gorm.Model{ID: 4}},
	//},)
	//fmt.Println(err.Error)
	// 多对多查询
	//var user1 User
	//if db.Preload("Languages").First(&user1, "id = ?", 2).RecordNotFound() {
	//	fmt.Println("sb")
	//	return
	//}
	//fmt.Println(user1)
}
