package main

import (
	"SUPPORT/config"
	"SUPPORT/router"
	// "fmt"
)

func main() {
	config.InitConfig()       //读取配置文件
	r := router.SetupRouter() //初始化路由
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080" //默认端口
	}

	r.Run(port)
}
