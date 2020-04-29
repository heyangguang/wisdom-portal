package models

import (
	"github.com/casbin/casbin/v2"
	"log"
	wisdomPortal "wisdom-portal/wisdom-portal"
)

type CasbinModel struct {
	//ID       int    `db:"id" json:"id"`
	Ptype    string `db:"p_type" json:"p_type"`
	RoleName string `db:"v0" json:"role_name"`
	Path     string `db:"v1" json:"path"`
	Method   string `db:"v2" json:"method"`
}

// 添加权限
func (c *CasbinModel) AddPerm(cm CasbinModel) bool {
	e := LoadPolicyPerm()
	ok, err := e.AddPolicy(cm.RoleName, cm.Path, cm.Method)
	if err != nil {
		log.Fatalf("casbin AddPolicy Failed, err:%v \n", err)
		return false
	}
	return ok
}

// 刷新策略到内存
// 重新加载数据到内存里，所以需要在上面AddPerm调用以后刷新
func LoadPolicyPerm() (e *casbin.Enforcer) {
	// 这里有一个坑，是MySQL和Golang的 nil != null 导致的，所以在数据库里给个默认值字符串的空，而不是null
	// TODB: 后面可以重写SqlxAdapter来解决这个问题
	e, err := casbin.NewEnforcer(wisdomPortal.BaseDir()+"/wisdom-portal/conf/rbac_model.conf", GormAdapter)
	if err != nil {
		log.Fatalf("casbin NewEnforcer Failed, err:%v \n", err)
		return
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatalf("casbin LoadPolicy Failed, err:%v \n", err)
		return
	}
	return
}
