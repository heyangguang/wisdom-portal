package user

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/register", register)
	routeGroup.GET("/user", test)
}
