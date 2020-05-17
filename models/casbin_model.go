package models

import (
	"errors"
	"github.com/casbin/casbin"
	"log"
	"time"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

// 权限模板
type Role struct {
	BaseModel
	RoleName    string   `gorm:"not null;comment:'权限名'" json:"role_name" binding:"required"`
	Remark      string   `gorm:"type:text;comment:'权限说明'" json:"remark"`
	RoleObjActs []ObjAct `gorm:"many2many:role_obj_acts" json:"roleObjActs"`
}

// 权限对象
type ObjAct struct {
	BaseModel
	ObjName string `gorm:"not null;comment:'对象'" json:"obj_name"`
	ActName string `gorm:"not null;comment:'动作'" json:"act_name"`
	Tag     string `gorm:"not null;comment:'权限分类'" json:"tag"`
}

// casbin权限模型
type CasbinRule struct {
	BaseModel
	PType string `gorm:"not null;comment:'权限类型'" json:"p_type"`
	Sub   string `gorm:"column:v0" json:"sub"`
	Obj   string `gorm:"column:v1" json:"obj"`
	Act   string `gorm:"column:v2" json:"act"`
	V3    string `json:"v3"`
	V4    string `json:"v4"`
	V5    string `json:"v5"`
}

// 权限详细条目
type RoleLine struct {
	ObJName string `json:"obj_name"`
	ActName string `json:"act_name"`
}

// 用户添加权限模板结构体
type AddPermUser struct {
	RoleId string `json:"role_id"`
}

// 用户组添加权限模板结构体
type AddPermUserGroup struct {
	RoleId string `json:"role_id"`
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
	var role Role
	logger.Debug("AddPermUserGroup")
	if err := DB.Where("id = ?", gid).First(&userGroup).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return err
	}
	if isOk := DB.Where("v0 = ?", userGroup.GroupName).First(&CasbinRule{}).RecordNotFound(); !isOk {
		return errors.New("the user group already has permission templates")
	}
	logger.Debug("用户组名" + userGroup.GroupName)
	if err := DB.Preload("RoleObjActs").Where("id = ?", a.RoleId).First(&role).Error; err != nil {
		logger.Error("AddPermUserGroup    " + err.Error())
		return err
	}
	logger.Debug("权限模板名" + role.RoleName)
	e := LoadPolicyPerm()
	for _, value := range role.RoleObjActs {
		isOk := e.AddPolicy(userGroup.GroupName, value.ObjName, value.ActName)
		if !isOk {
			logger.Error("AddPermUserGroup    " + "the current usergroup already has this permission")
			return errors.New("the current usergroup already has this permission")
		}
	}
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
			for _, roleObj := range role.RoleObjActs {
				DB.Where("p_type = ? and v0 = ? and v1 = ? and v2 = ?", "p", userGroup.GroupName, roleObj.ObjName, roleObj.ActName).Delete(&CasbinRule{})
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
	return nil
}

// 用户添加权限模板
func (a *AddPermUser) AddPermUser(uid string) error {
	var user User
	var role Role
	logger.Debug("AddPermUser开始")
	if err := DB.Where("id = ?", uid).First(&user).Error; err != nil {
		logger.Error("AddPermUser    " + err.Error())
		return err
	}
	logger.Debug("用户名" + user.UserName)
	if err := DB.Preload("RoleObjActs").Where("id = ?", a.RoleId).First(&role).Error; err != nil {
		logger.Error("AddPermUser    " + err.Error())
		return err
	}
	logger.Debug("权限模板名" + role.RoleName)
	e := LoadPolicyPerm()
	for _, value := range role.RoleObjActs {
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
	return nil
}

// 添加权限模板
func (c *Role) AddPerm(role Role) error {
	var r Role
	// 创建权限模板
	if err := DB.Omit("id").Set("gorm:save_associations", false).Create(&role).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	// 获取创建的权限模板对象
	if err := DB.Select("id").Where("role_name = ?", role.RoleName).First(&r).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	// 判断权限对象是否存在
	for _, value := range role.RoleObjActs {
		if isOk := DB.Where("id = ?", value.ID).First(&ObjAct{}).RecordNotFound(); isOk {
			DB.Delete(&r)
			logger.Error("AddPerm    选择的权限对象不存在")
			return errors.New("选择的权限对象不存在")
		}
	}
	// 添加权限对象
	if err := DB.Model(&r).Association("RoleObjActs").Append(&role.RoleObjActs).Error; err != nil {
		logger.Error("AddPerm    " + err.Error())
		return err
	}
	return nil
}

// 刷新策略到内存
// 重新加载数据到内存里，所以需要在上面AddPerm调用以后刷新
func LoadPolicyPerm() (e *casbin.Enforcer) {
	// 这里有一个坑，是MySQL和Golang的 nil != null 导致的，所以在数据库里给个默认值字符串的空，而不是null
	// TODB: 后面可以重写SqlxAdapter来解决这个问题
	e = casbin.NewEnforcer(wisdomPortal.BaseDir()+"/wisdom-portal/conf/rbac_model.conf", GormAdapter)
	err := e.LoadPolicy()
	if err != nil {
		log.Fatalf("casbin LoadPolicy Failed, err:%v \n", err)
		return
	}
	return
}
