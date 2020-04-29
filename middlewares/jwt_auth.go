package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal/logger"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthCheckToken(skipper ...SkipperFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 如果skip，就不需要执行下面逻辑了
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		logger.Debug("需要验证登录")
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token验证失败，头部请求为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token验证失败，格式错误",
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，解析
		mc, err := models.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token验证失败，无效的token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
