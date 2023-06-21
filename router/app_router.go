package router

import (
	"ginchat/service"
	"ginchat/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var mode string

var Log = &util.MyLog{}

func init() {
	Log.Debugln("start gin server")
	mode = viper.GetString("env")
	Log.Debugf("start set mode:  %s\n", mode)
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	Log.Debugln("end set mode")
	Log.Debugln("success start")
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
