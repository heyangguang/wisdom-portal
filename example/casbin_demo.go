package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	wisdomPortal "wisdom-portal/wisdom-portal"
)

func main() {

	a := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test", true)
	e := casbin.NewEnforcer(wisdomPortal.BaseDir()+"/wisdom-portal/conf/rbac_model.conf", a)
	//从DB加载策略
	e.LoadPolicy()

	//获取router路由对象
	r := gin.New()
	//使用自定义拦截器中间件
	r.Use(LanjieqiHandler(e))
	//创建请求
	r.GET("/api/v1/aa", func(c *gin.Context) {
		fmt.Println("Hello 接收到GET请求..")
	})

	r.Run(":9090") //参数为空 默认监听8080端口
}

//拦截器
func LanjieqiHandler(e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"

		//判断策略中是否存在
		if e.Enforce(sub, obj, act) {
			fmt.Println("通过权限")
			c.Next()
		} else {
			fmt.Println("权限没有通过")
			c.Abort()
		}
	}
}
