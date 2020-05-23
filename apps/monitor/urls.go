package monitor

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/monitor", getMonitor)
	routeGroup.POST("/monitor", addMonitor)
}
