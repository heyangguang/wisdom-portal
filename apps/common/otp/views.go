package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
)

// 二维码
func createQrCode(c *gin.Context)  {
	qrCodeData := c.DefaultQuery("data", "")
	secret := c.DefaultQuery("secret", "")
	imgPath := wisdomPortal.BaseDir() + "/static/"
	logger.Debug("二维码生成" + qrCodeData)
	logger.Debug("二维码生成" + secret)
	if qrCodeData != "" && secret != "" {
		err := qrcode.WriteFile(qrCodeData, qrcode.Medium, 256, imgPath + "qr_" + secret + ".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "msg": "二维码生成失败请联系管理员"})
			logger.Info("createQrCode" + err.Error())
			return
		}
		logger.Debug("二维码生成" + c.Request.RequestURI)
		c.JSON(http.StatusOK, gin.H{"code": 1,  "data": "http://127.0.0.1:8080/api/v1/view-qr-code/?imagename=" + "qr_" + secret + ".png"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
	return
}

// 显示二维码
func viewQrCode(c *gin.Context)  {
	imageName := c.Query("imagename")
	filePath := wisdomPortal.BaseDir() + "/static/"
	file, err := ioutil.ReadFile(filePath + imageName)
	if err != nil {
		logger.Debug("viewQrCode" + err.Error())
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "二维码访问失败"})
	}
	c.Writer.WriteString(string(file))
	return
}
