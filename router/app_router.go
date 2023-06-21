package router

import (
	"ginchat/service"
	"ginchat/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var mode string

func init() {
	util.Logger.Println("start gin server")
	mode = viper.GetString("env")
	util.Logger.Printf("start set mode:  %s\n", mode)
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	util.Logger.Println("end set mode")
	util.Logger.Println("success start")
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
