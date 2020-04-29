package jwt

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/login", login)
}
