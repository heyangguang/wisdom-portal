package monitor

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/monitor/tcp", getTcpMonitor)
	routeGroup.GET("/monitor/tcp/quality", getTcpMonitorQuality)
	routeGroup.GET("/monitor/intermediate", queryIntermediateMonitor)
	routeGroup.GET("/monitor/access", queryAccessLogMonitor)
	routeGroup.POST("/monitor/tcp", addMonitor)
	routeGroup.POST("/monitor/intermediate", addIntermediateMonitor)
	routeGroup.POST("/monitor/access", addAccessLogMonitor)
}
