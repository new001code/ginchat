package router

import (
	"ginchat/common"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var mode string

func init() {
	gin.ForceConsoleColor()
	common.Logger.Info("start gin server")
	mode = viper.GetString("env")
	common.Logger.Infof("start set mode:  %s", mode)
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	common.Logger.Info("end set mode")
	common.Logger.Info("success start")
}

func Router() *gin.Engine {
	r := gin.Default()

	//user
	userGroup := r.Group("/user")
	{
		userService := &service.UserService{}
		userGroup.POST("/register", userService.Register)
		userGroup.POST("/checkCode", userService.LoginCheckCode)
	}
	return r
}
