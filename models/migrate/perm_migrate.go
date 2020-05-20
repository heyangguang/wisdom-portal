package migrate

import "wisdom-portal/models"

// 权限模板
type Rule struct {
	models.BaseMigrate
	RuleName    string   `gorm:"not null;comment:'权限名'" json:"rule_name" binding:"required"`
	Remark      string   `gorm:"type:text;comment:'权限说明'" json:"remark"`
	RuleObjActs []ObjAct `gorm:"many2many:rule_obj_acts" json:"ruleObjActs"`
}

// 权限对象
type ObjAct struct {
	models.BaseMigrate
	ObjName string `gorm:"not null;comment:'对象'" json:"obj_name"`
	ActName string `gorm:"not null;comment:'动作'" json:"act_name"`
	Tag     string `gorm:"not null;comment:'权限分类'" json:"tag"`
}

// casbin权限模型
type CasbinRule struct {
	models.BaseMigrate
	PType string `gorm:"not null;comment:'权限类型'" json:"p_type"`
	Sub   string `gorm:"column:v0" json:"sub"`
	Obj   string `gorm:"column:v1" json:"obj"`
	Act   string `gorm:"column:v2" json:"act"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}
