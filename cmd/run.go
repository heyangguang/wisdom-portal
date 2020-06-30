/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"time"
	"wisdom-portal/apps/alert"
	"wisdom-portal/apps/common/health"
	"wisdom-portal/apps/common/jwt"
	"wisdom-portal/apps/common/otp"
	"wisdom-portal/apps/common/permission"
	"wisdom-portal/apps/common/user"
	"wisdom-portal/apps/common/usergroup"
	"wisdom-portal/apps/monitor"
	"wisdom-portal/models"
	wisdom_portal "wisdom-portal/wisdom-portal"
	"wisdom-portal/wisdom-portal/clear_static"
	"wisdom-portal/wisdom-portal/forms"
	"wisdom-portal/wisdom-portal/logger"
	v1 "wisdom-portal/wisdom-portal/routers/api/v1"
)

var port int
var log string
var logLevel string
var db string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start wisdomServer port default 8080",
	Run: func(cmd *cobra.Command, args []string) {
		if port == 0 || log == "" || logLevel == "" || db == "" {
			_ = cmd.Help()
			return
		} else {
			runStart(log, logLevel, db, port)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&log, "log", "l", "./logs/wisdom-client.log", "--log <logLocation> example --config ./logs/wisdom-portal.log")
	runCmd.Flags().IntVarP(&port, "port", "p", 8080, "--port <port> example --port 8080")
	runCmd.Flags().StringVar(&logLevel, "logLevel", "DEBUG", "--logLevel <logLocation> example --logLevel DEBUG")
	runCmd.Flags().StringVar(&db, "db", "", "--db <logLocation> example --db root:123456@(127.0.0.1:3306)/test")
}

func runStart(logPath, logLevel, db string, port int) {
	// 初始化日志模块
	//logPath := wisdom_portal.BaseDir() + "/logs/wisdom-portal.log"
	err := logger.InitLogger(logPath, 1, 7, 10, logLevel)
	if err != nil {
		fmt.Println(err.Error())
	}

	// !DEBUG
	if logLevel != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		logger.Info(fmt.Sprintf("Listening and serving HTTP on :%d", port))
	}

	// 加载数据库
	models.DBConnectInit(db)

	// 加载表单验证模块
	forms.InitFormValidate()

	// 注册自定义表单方法
	forms.RegisterCustomValidationFunc(user.CustomValidations, usergroup.CustomValidations, monitor.CustomValidations)

	// 初始化自定义表单方法
	forms.InitCustomValidationFunc()

	// 加载多个APP的路由配置
	v1.Include(permission.Routers, user.Routers, jwt.Routers,
		otp.Routers, usergroup.Routers, monitor.Routers, alert.Routers, health.Routers)

	// 加载清理static异步模块
	fileChan := make(chan string, 50)
	go clear_static.GetAllFile(wisdom_portal.BaseDir()+"/static/", fileChan, time.Second*60*10)
	for i := 0; i <= 5; i++ {
		go clear_static.CheckFileDiffTime(fileChan)
	}

	// 初始化路由
	r := v1.InitV1()
	_ = r.Run(fmt.Sprintf(":%d", port))
}
