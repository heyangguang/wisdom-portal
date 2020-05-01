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
// @Success 200 {object} gin.H "{"code": 10000, "msg": "成功", "data": ""}"
// @Failure 500 {object} gin.H "{"code": 60001, "msg": "内部系统接口调用异常", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效"}"
// @Router /api/v1/create-qr-code [GET]
func createQrCode(c *gin.Context) {
	qrCodeData := c.DefaultQuery("data", "")
	secret := c.DefaultQuery("secret", "")
	imgPath := wisdomPortal.BaseDir() + "/static/"
	logger.Debug("二维码生成" + qrCodeData)
	logger.Debug("二维码生成" + secret)
	if qrCodeData != "" && secret != "" {
		err := qrcode.WriteFile(qrCodeData, qrcode.Medium, 256, imgPath+"qr_"+secret+".png")
		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{
					"code": result.InterfaceInnerInvokeError,
					"msg":  result.ResultText(result.InterfaceInnerInvokeError),
					"err":  err.Error(),
				})
			logger.Info("createQrCode" + err.Error())
			return
		}
		logger.Debug("二维码生成" + c.Request.RequestURI)
		c.JSON(http.StatusOK, gin.H{
			"code": result.SuccessCode,
			"msg":  result.ResultText(result.SuccessCode),
			"data": "http://127.0.0.1:8080/api/v1/view-qr-code/?imagename=" + "qr_" + secret + ".png",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"code": result.ParamInvalid, "msg": result.ResultText(result.ParamInvalid)})
	return
}

// @Summary 显示多因子认证二维码
// @Description 用于用户绑定多因子
// @Tags 多因子
// @accept json
// @Param imagename query string true "图片名"
// @Success 200
// @Failure 500 {object} gin.H "{"code": 60001, "msg": "内部系统接口调用异常", "err": ""}"
// @Failure 400 {object} gin.H "{"code": 10001, "msg": "参数无效"}"
// @Failure 404 {object} gin.H "{"code": 50001, "msg": "数据未找到", "err": ""}"
// @Router /api/v1/view-qr-code [GET]
func viewQrCode(c *gin.Context) {
	imageName := c.DefaultQuery("imagename", "")
	if imageName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": result.ParamInvalid,
			"msg":  result.ResultText(result.ParamInvalid),
		})
	}
	filePath := wisdomPortal.BaseDir() + "/static/"
	file, err := ioutil.ReadFile(filePath + imageName)
	if err != nil {
		logger.Error("viewQrCode" + err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"code": result.DataNone,
			"msg":  result.ResultText(result.DataNone),
			"err":  err.Error(),
		})
	}
	if _, err := c.Writer.WriteString(string(file)); err != nil {
		logger.Error("viewQrCode" + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": result.InterfaceInnerInvokeError,
			"msg":  result.ResultText(result.InterfaceInnerInvokeError),
			"err":  err.Error(),
		})
	}
	return
}
