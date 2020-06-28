package monitor

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/monitor/tcp", getTcpMonitor)
	routeGroup.GET("/monitor/tcp/quality", getTcpMonitorQuality)
	routeGroup.GET("/monitor/intermediate", queryIntermediateMonitor)
	routeGroup.POST("/monitor/tcp", addMonitor)
	routeGroup.POST("/monitor/intermediate", addIntermediateMonitor)
}
