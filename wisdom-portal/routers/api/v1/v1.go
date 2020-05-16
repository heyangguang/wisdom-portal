package v1

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "wisdom-portal/docs"
	"wisdom-portal/middlewares"
	"wisdom-portal/wisdom-portal/logger"
)

type Option func(routeGroup *gin.RouterGroup)

var options []Option

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func InitV1() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	engine.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	engine.Use(middlewares.Cors())
	v1Group := engine.Group("api/v1/")
	// use ginSwagger middleware to
	v1Group.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Group.Use(middlewares.JWTAuthCheckToken(
		middlewares.AllowMethodAndPathPrefixSkipper(
			middlewares.JoinRouter("POST", "/api/v1/login"),
			middlewares.JoinRouter("POST", "/api/v1/register"),
			middlewares.JoinRouter("GET", "/api/v1/view-qr-code"),
			middlewares.JoinRouter("GET", "/api/v1/create-qr-code"),
		),
	))
	v1Group.Use(middlewares.PermAuthCheckRole(
		middlewares.AllowMethodAndPathPrefixSkipper(
			middlewares.JoinRouter("POST", "/api/v1/login"),
			middlewares.JoinRouter("POST", "/api/v1/register"),
			middlewares.JoinRouter("GET", "/api/v1/view-qr-code"),
			middlewares.JoinRouter("GET", "/api/v1/create-qr-code"),
			middlewares.JoinRouter("GET", "/api/v1/pub/current/user"),
		),
	))
	for _, opt := range options {
		opt(v1Group)
	}
	return engine
}
