package permission

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/perm", test)
	routeGroup.POST("/perm", addPerm)
	// 用户权限
	routeGroup.POST("/perm/user/:uid", addPermUser)
}
