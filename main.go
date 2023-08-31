package main

import (
	"fmt"
	"user-api/global"
	"user-api/initialize"
	"user-api/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	// 1. 初始化配置文件
	initialize.InitConfig()
	// 2. 初始化日志Logger
	initialize.InitLogger()
	// 3. 初始化Router
	router := initialize.Routers()
	err := router.Run(fmt.Sprintf(":%d", global.Conf.App.Port))
	// global.Log.Info("Start server! Listen: 0.0.0.0:", global.Conf.App.Port)
	if err != nil {
		panic(err)
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", validators.Mobile)
	}
}
