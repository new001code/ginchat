package router

import (
	"ginchat/service"
	"ginchat/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var mode string

func init() {
	util.DebugLogger.Println("start gin server")
	mode = viper.GetString("env")
	util.DebugLogger.Printf("start set mode:  %s\n", mode)
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	util.DebugLogger.Println("end set mode")
	util.DebugLogger.Println("success start")
}

func Router() *gin.Engine {
	r := gin.Default()

	//user
	userGroup := r.Group("/user")
	{
		userService := &service.UserService{}
		userGroup.POST("/register", userService.Register)
	}
	return r
}
