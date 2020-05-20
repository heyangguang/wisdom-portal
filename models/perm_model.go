package models

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin"
	"log"
	"time"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

// 权限模型
type CasbinRule struct {
	BaseModel
	PType string `json:"p_type"`
	Sub   string `json:"sub"`
	Obj   string `json:"obj"`
	Act   string `json:"act"`
}

// 权限模板
type Rule struct {
	BaseModel
	RuleName    string   `json:"rule_name" validate:"required" label:"rule_name"`
	Remark      string   `json:"remark"`
	RuleObjActs []ObjAct `gorm:"many2many:rule_obj_acts" json:"rule_obj_acts" validate:"required" label:"rule_obj_acts"`
}

// 权限对象
type ObjAct struct {
	ID      int    `json:"id"`
	ObjName string `json:"-"`
	ActName string `json:"-"`
	Tag     string `json:"-"`
}

// 权限详细条目
type RuleLine struct {
	ObJName string `json:"obj_name"`
	ActName string `json:"act_name"`
}

// 用户添加权限模板结构体
type AddPermUser struct {
	RuleId int `json:"rule_id" validate:"required" label:"rule_id"`
}

// 用户组添加权限模板结构体
type AddPermUserGroup struct {
	RuleId int `json:"rule_id" validate:"required" label:"rule_id"`
}

// 直接对权限实体表赋权
// 主要用于注册用户的时候需要给默认权限
func AddDefaultPerm(userName, objName, actName string) error {
	// 添加权限对象
	e := LoadPolicyPerm()
	isOk := e.AddPolicy(userName, objName, actName)
	if !isOk {
		logger.Error("AddDefaultPerm    " + "the current user already has this permission")
		return errors.New("the current user already has this permission")
	}
	// 更新时间戳
	data := make(map[string]interface{})
	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()
	DB.Model(CasbinRule{}).Where("p_type = ? and v0 = ?", "p", userName).Updates(data)
	return nil
}

// 用户组添加权限模板
func (a *AddPermUserGroup) AddPermUserGroup(gid string) error {
	var userGroup UserGroup
	var rule Rule
	logger.Debug("AddPermUserGroup")
	if err := DB.Where("id = ?", gid).First(&userGroup).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return err
	}

	logger.Debug("用户组名" + userGroup.GroupName)
	logger.Debug(fmt.Sprintf("用户组详情：%v", userGroup))

	logger.Debug("判断用户组是否已经有权限模板")
	if userGroup.RuleId != 0 {
		logger.Error("AddPermUserGroup    " + "the user group already has a permission template")
		return errors.New("the user group already has a permission template")
	}

	logger.Debug("拿权限模板的所有obj和act")
	if err := DB.Preload("RuleObjActs").Where("id = ?", a.RuleId).First(&rule).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return errors.New("not found rule record")
	}

	logger.Debug("权限模板名" + rule.RuleName)
	logger.Debug("同步用户组权限到权限表里")
	e := LoadPolicyPerm()
	for _, value := range rule.RuleObjActs {
		isOk := e.AddPolicy(userGroup.GroupName, value.ObjName, value.ActName)
		if !isOk {
			logger.Error("AddPermUserGroup    " + "the current usergroup already has this permission")
			return errors.New("the current user group already has this permission")
		}
	}

	logger.Debug("同步用户组下的用户权限到权限表里")
	// 权限表关联用户组下的用户
	// 1. 查询用户组下有哪些用户
	if err := DB.Preload("Users").Where("id = ?", gid).First(&userGroup).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return err
	}
	// 2. 遍历用户组下的用户，批量添加对应关系
	for _, value := range userGroup.Users {
		if isOk := e.AddGroupingPolicy(value.UserName, userGroup.GroupName); !isOk {
			logger.Error("AddPermUserGroup    " + "users who already have this user group under current permissions")
			for _, ruleObj := range rule.RuleObjActs {
				DB.Where("p_type = ? and v0 = ? and v1 = ? and v2 = ?", "p", userGroup.GroupName, ruleObj.ObjName, ruleObj.ActName).Delete(&CasbinRule{})
			}
			return errors.New("users who already have this user group under current permissions")
		}
	}

	// 3. 更新时间戳
	data := make(map[string]interface{})
	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()
	DB.Model(CasbinRule{}).Where("p_type = ? and v0 = ?", "p", userGroup.GroupName).Updates(data)
	DB.Model(CasbinRule{}).Where("p_type = ? and v1 = ?", "g", userGroup.GroupName).Updates(data)

	// 更新用户组表 role_id 字段
	if err := DB.Model(&userGroup).Update("rule_id", a.RuleId).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return err
	}
	return nil
}

// 用户添加权限模板
func (a *AddPermUser) AddPermUser(uid string) error {
	var user User
	var rule Rule
	logger.Debug("AddPermUser开始")
	if err := DB.Where("id = ?", uid).First(&user).Error; err != nil {
		logger.Error("AddPermUser    " + err.Error())
		return err
	}

	logger.Debug("用户名" + user.UserName)
	logger.Debug(fmt.Sprintf("详情信息：%v", user))

	logger.Debug("判断用户是否已经有权限模板")
	if user.RuleId != 0 {
		logger.Error("AddPermUser    " + "the user already has a permission template")
		return errors.New("the user already has a permission template")
	}

	logger.Debug("拿权限模板的所有obj和act")
	if err := DB.Preload("RuleObjActs").Where("id = ?", a.RuleId).First(&rule).Error; err != nil {
		logger.Error("AddPermUser    " + err.Error())
		return err
	}

	logger.Debug("权限模板名" + rule.RuleName)
	logger.Debug("同步权限到权限表")
	e := LoadPolicyPerm()
	for _, value := range rule.RuleObjActs {
		isOk := e.AddPolicy(user.UserName, value.ObjName, value.ActName)
		if !isOk {
			logger.Error("AddPermUser    " + "the current user already has this permission")
			return errors.New("the current user already has this permission")
		}
	}

	// 更新时间戳
	data := make(map[string]interface{})
	data["created_at"] = time.Now()
	data["updated_at"] = time.Now()
	DB.Model(CasbinRule{}).Where("p_type = ? and v0 = ?", "p", user.UserName).Updates(data)

	// 更新用户表 role_id 字段
	if err := DB.Model(&user).Update("rule_id", a.RuleId).Error; err != nil {
		logger.Error("AddPermUser    " + err.Error())
		return err
	}
	return nil
}

// 添加权限模板
func (t *Rule) AddPerm(rule Rule) error {
	var r Rule
	// 创建权限模板
	if err := DB.Omit("id").Set("gorm:save_associations", false).Create(&rule).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	// 获取创建的权限模板对象
	if err := DB.Select("id").Where("rule_name = ?", rule.RuleName).First(&r).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	// 判断权限对象是否存在
	logger.Debug(fmt.Sprintf("%v", rule.RuleObjActs))
	for _, value := range rule.RuleObjActs {
		logger.Debug(fmt.Sprintf("%d", value.ID))
		if isOk := DB.Where("id = ?", value.ID).First(&ObjAct{}).RecordNotFound(); isOk {
			DB.Delete(&r)
			logger.Error("AddPerm    选择的权限对象不存在")
			return errors.New("选择的权限对象不存在")
		}
	}
	// 添加权限对象
	if err := DB.Model(&r).Association("RuleObjActs").Append(&rule.RuleObjActs).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	return nil
}

// 刷新策略到内存
// 重新加载数据到内存里，所以需要在上面AddPerm调用以后刷新
func LoadPolicyPerm() (e *casbin.Enforcer) {
	// 这里有一个坑，是MySQL和Golang的 nil != null 导致的，所以在数据库里给个默认值字符串的空，而不是null
	e = casbin.NewEnforcer(wisdomPortal.BaseDir()+"/wisdom-portal/conf/rbac_model.conf", GormAdapter)
	err := e.LoadPolicy()
	if err != nil {
		log.Fatalf("casbin LoadPolicy Failed, err:%v \n", err)
		return
	}
	return
}
