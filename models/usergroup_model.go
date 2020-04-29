package models

type UserGroup struct {
	BaseModel
	GroupName string `gorm:"comment:'用户组名'"`
	Remark    string `gorm:"type:text;comment:'备注'"`
	Users     []User `gorm:"many2many:user_usergroup"`
}
