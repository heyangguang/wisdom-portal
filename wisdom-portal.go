package main

import "wisdom-portal/cmd"

// @title WisdomPortal backend API
// @version v1.0
// @description This is a wisdomPortal backend interface system server.

// @contact.name API Support
// @contact.url http://mail.csic711.com/
// @contact.email heyangev@cn.ibm.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.31.2:8080

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /

func main() {
	// initCmd
	cmd.Execute()
}
