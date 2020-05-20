package migrate

import "wisdom-portal/models"

type UserGroupMigrate struct {
	models.BaseMigrate
	GroupName string        `gorm:"unique;comment:'用户组名';not null" binding:"required" json:"group_name"`
	Remark    string        `gorm:"type:text;comment:'备注'" json:"remark"`
	RuleId    int           `gorm:"default:null" json:"rule_id"`
	Rule      RuleMigrate   `gorm:"foreignkey:RuleId;comment:'权限模板'" json:"rule"`
	Users     []UserMigrate `gorm:"many2many:user_usergroup;association_jointable_foreignkey:user_id;jointable_foreignkey:user_group_id" json:"users"`
}

func (t *UserGroupMigrate) TableName() string {
	return "user_group"
}
