package alert

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.POST("/alert", addAlert)
	routeGroup.PUT("/alert/:id", updateAlert)
}
