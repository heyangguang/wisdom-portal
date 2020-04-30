package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/models"
)

// 权限检查中间件
func PermAuthCheckRole(skipper ...SkipperFunc) gin.HandlerFunc {
	// TODB 正常来说应该使用JWT登录后，获取权限，这里还没有实现JWT，简单写一个直接判断权限，来测试。
	return func(c *gin.Context) {
		// 如果skip，就不需要执行下面逻辑了
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		var username interface{}
		username, isExist := c.Get("username")
		if isExist {
			//name := c.Query("name")
			e := models.LoadPolicyPerm()
			// 获取用户和用户组的全部权限
			userRoles := e.GetImplicitPermissionsForUser(username.(string))
			fmt.Println(userRoles)
			// 检查权限
			for _, value := range userRoles {
				if util.KeyMatch(c.Request.URL.Path, value[1]) && (value[2] == c.Request.Method || value[2] == "*") {
					isOk := e.Enforce(value[0], value[1], value[2])
					if isOk {
						c.Next()
						return
					}
				}
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"status": -1,
			"msg":    "很抱歉您没有此权限",
		})
		c.Abort()
		return
	}
}
