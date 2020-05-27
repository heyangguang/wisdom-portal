package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"wisdom-portal/wisdom-portal/logger"
)

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，白名单
func AllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(path)
		logger.Debug("AllowMethodAndPathPrefixSkipper中间件    " + path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// NoAllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，黑名单
func NoAllowMethodAndPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := JoinRouter(c.Request.Method, c.Request.URL.Path)
		pathLen := len(path)
		logger.Debug("NoAllowMethodAndPathPrefixSkipper中间件    " + path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}

		return true
	}
}

// JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
