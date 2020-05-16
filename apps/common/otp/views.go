package otp

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	wisdomPortal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
	"wisdom-portal/wisdom-portal/result"
)

// @Summary 生成多因子认证二维码
// @Description 用于用户绑定多因子
// @Tags 多因子
// @accept json
// @Produce  json
// @Param data query string true "数据"
// @Param secret query string true "秘钥"
// @Success 200 {object} result.CreateQrCodeResult "{"code": 10000}"
// @Failure 500 {object} result.FailResult "{"code": 60001}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Router /api/v1/create-qr-code [GET]
func createQrCode(c *gin.Context) {
	qrCodeData := c.DefaultQuery("data", "")
	secret := c.DefaultQuery("secret", "")
	imgPath := wisdomPortal.BaseDir() + "/static/"
	logger.Debug("二维码生成" + qrCodeData)
	logger.Debug("二维码生成" + secret)
	if qrCodeData != "" && secret != "" {
		// TODO: 这里的Secret应该去数据库里验证是否存在，否则存在无限通过接口生成二维码文件的漏洞
		err := qrcode.WriteFile(qrCodeData, qrcode.Medium, 256, imgPath+"qr_"+secret+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError, result.NewFailResult(result.InterfaceInnerInvokeError, err.Error()))
			logger.Info("createQrCode" + err.Error())
			return
		}
		logger.Debug("二维码生成" + c.Request.RequestURI)
		resultData := "/api/v1/view-qr-code/?imagename=" + "qr_" + secret + ".png"
		c.JSON(http.StatusOK, result.NewCreateQrCodeResult(result.SuccessCode, resultData))
		return
	}
	c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, "data or secret filed input error"))
	return
}

// @Summary 显示多因子认证二维码
// @Description 用于用户绑定多因子
// @Tags 多因子
// @accept json
// @Param imagename query string true "图片名"
// @Success 200
// @Failure 500 {object} result.FailResult "{"code": 60001}"
// @Failure 400 {object} result.FailResult "{"code": 10001}"
// @Failure 404 {object} result.FailResult "{"code": 50001}"
// @Router /api/v1/view-qr-code [GET]
func viewQrCode(c *gin.Context) {
	imageName := c.DefaultQuery("imagename", "")
	if imageName == "" {
		c.JSON(http.StatusBadRequest, result.NewFailResult(result.ParamInvalid, "imageName filed input error"))
	}
	filePath := wisdomPortal.BaseDir() + "/static/"
	file, err := ioutil.ReadFile(filePath + imageName)
	if err != nil {
		logger.Error("viewQrCode" + err.Error())
		c.JSON(http.StatusNotFound, result.NewFailResult(result.DataNone, err.Error()))
	}
	if _, err := c.Writer.WriteString(string(file)); err != nil {
		logger.Error("viewQrCode" + err.Error())
		c.JSON(http.StatusInternalServerError, result.NewFailResult(result.InterfaceInnerInvokeError, err.Error()))
	}
	return
}
