package user

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/register", register)
	routeGroup.GET("/pub/current/user", getCurrentUser)
}
