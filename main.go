package main

import (
	"fmt"
	"wisdom-portal/apps/common/jwt"
	"wisdom-portal/apps/common/otp"
	"wisdom-portal/apps/common/permission"
	"wisdom-portal/apps/common/user"
	"wisdom-portal/models"
	"wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/logger"
	v1 "wisdom-portal/wisdom-portal/routers/api/v1"
)

// @title Swagger Example API
// @version v1.6.5
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
func main() {
	// 初始化日志模块
	logPath := wisdom_portal.BaseDir() + "/logs/wisdom-portal.log"
	err := logger.InitLogger(logPath, 1, 7, 10, "DEBUG")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 加载数据库
	models.DBConnectInit()

	// 加载多个APP的路由配置
	v1.Include(permission.Routers, user.Routers, jwt.Routers, otp.Routers)

	// 初始化路由
	r := v1.InitV1()
	_ = r.Run(":8080")
}
