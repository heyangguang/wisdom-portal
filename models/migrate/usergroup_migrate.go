package migrate

import "wisdom-portal/models"

type UserGroupMigrate struct {
	models.BaseMigrate
	GroupName string        `gorm:"unique;comment:'用户组名';not null" binding:"required" json:"group_name"`
	Remark    string        `gorm:"type:text;comment:'备注'" json:"remark"`
	Users     []UserMigrate `gorm:"many2many:user_usergroup" json:"users"`
}
