package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wisdom-portal/wisdom-portal/result"
)

// 验证客户端的中间件
func ClientAuth(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		token := c.Query("token")
		if token == "AeZmOUaO-oVlRmeGtHK539xXt4nfh6DSCnVArRq-41M" {
			c.Next()
			return
		}
		c.Abort()
		c.JSON(http.StatusForbidden, result.NewSuccessResult(result.PermissionNoAccess))
		return
	}
}
