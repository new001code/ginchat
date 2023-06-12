package router

import (
	"ginchat/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var mode string

func init() {
	log.Println("start gin server")
	mode = viper.GetString("env")
	log.Printf("start set mode:  %s\n", mode)
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Println("end set mode")
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
