package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/wisdom-portal/result"
)

// 用于检测服务是否正常
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, result.NewSuccessResult(result.SuccessCode))
}
