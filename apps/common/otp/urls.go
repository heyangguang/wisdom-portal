package otp

import "github.com/gin-gonic/gin"

func Routers(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/create-qr-code", createQrCode)
	routeGroup.GET("/view-qr-code", viewQrCode)
}

