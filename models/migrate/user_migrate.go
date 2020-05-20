package migrate

import "wisdom-portal/models"

type UserMigrate struct {
	models.BaseMigrate
	Name       string             `gorm:"unique;not null;comment:'姓名'" json:"name" binding:"required"`
	UserName   string             `gorm:"unique;not null;comment:'用户名' "json:"user_name" binding:"required"`
	PassWord   string             `gorm:"not null;comment:'密码'" json:"pass_word" binding:"required"`
	Phone      string             `gorm:"comment:'手机号'" json:"phone"`
	Email      string             `gorm:"comment:'电子邮箱'" json:"email"`
	Status     bool               `gorm:"comment:'用户状态'" json:"status"`
	Secret     string             `gorm:"comment:'双因子秘钥'" json:"secret"`
	UserGroups []UserGroupMigrate `gorm:"many2many:user_usergroup" json:"user_groups"`
}
