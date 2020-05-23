package migrate

import "wisdom-portal/models"

// 权限模板
type RuleMigrate struct {
	models.BaseMigrate
	RuleName    string          `gorm:"not null;comment:'权限名'" json:"rule_name"`
	Remark      string          `gorm:"type:text;comment:'权限说明'" json:"remark"`
	RuleObjActs []ObjActMigrate `gorm:"many2many:rule_obj_acts;association_jointable_foreignkey:obj_act_id;jointable_foreignkey:rule_id" json:"ruleObjActs"`
}

// 权限对象
type ObjActMigrate struct {
	models.BaseMigrate
	ObjName string `gorm:"not null;comment:'对象'" json:"obj_name"`
	ActName string `gorm:"not null;comment:'动作'" json:"act_name"`
	Tag     string `gorm:"not null;comment:'权限分类'" json:"tag"`
}

// casbin权限模型
type CasbinRuleMigrate struct {
	models.BaseMigrate
	PType string `gorm:"not null;comment:'权限类型'" json:"p_type"`
	Sub   string `gorm:"column:v0" json:"sub"`
	Obj   string `gorm:"column:v1" json:"obj"`
	Act   string `gorm:"column:v2" json:"act"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

func (t *RuleMigrate) TableName() string {
	return "rule"
}

func (t *ObjActMigrate) TableName() string {
	return "obj_act"
}

func (t *CasbinRuleMigrate) TableName() string {
	return "casbin_rule"
}
