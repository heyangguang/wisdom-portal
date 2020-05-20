package migrate

import "wisdom-portal/models"

type UserMigrate struct {
	models.BaseMigrate
	Name       string             `gorm:"not null;comment:'姓名'" json:"name" binding:"required"`
	UserName   string             `gorm:"unique;not null;comment:'用户名' "json:"user_name" binding:"required"`
	PassWord   string             `gorm:"not null;comment:'密码'" json:"pass_word" binding:"required"`
	Phone      string             `gorm:"comment:'手机号'" json:"phone"`
	Email      string             `gorm:"comment:'电子邮箱'" json:"email"`
	Status     bool               `gorm:"comment:'用户状态'" json:"status"`
	Secret     string             `gorm:"comment:'双因子秘钥'" json:"secret"`
	RuleId     int                `gorm:"default:null;comment:'权限模板'" json:"rule_id"`
	Rule       RuleMigrate        `gorm:"ForeignKey:RuleId" json:"rule_id"`
	UserGroups []UserGroupMigrate `gorm:"many2many:user_usergroup;association_jointable_foreignkey:user_group_id;jointable_foreignkey:user_id" json:"user_groups"`
}

func (t *UserMigrate) TableName() string {
	return "user"
}
