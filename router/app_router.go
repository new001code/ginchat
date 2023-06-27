package router

import (
	"fmt"
	"ginchat/common"
	"ginchat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	mode string
)

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
	gin.ForceConsoleColor()

	r := gin.New()
	r.Use(ginLogger())
	r.Use(cors())
	r.Use(gin.Recovery())

	//user
	userGroup := r.Group("/user")
	{
		userService := &service.UserService{}
		userGroup.POST("/register", userService.Register)
		userGroup.POST("/checkCode", userService.LoginCheckCode)
		userGroup.GET("/publicKey", userService.GetPublicKey)
	}

	return r
}

func ginLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		statusColor := param.StatusCodeColor()
		methodColor := param.MethodColor()
		resetColor := param.ResetColor()

		return fmt.Sprintf("%s\t%s%s%s\t[%s]\t%s\t%s|%s %d %s|\t%s\t%s\n",
			param.TimeStamp.Format("2006-01-02 15:04:05.000000"),
			methodColor,
			param.Method,
			resetColor,
			param.ClientIP,
			param.Path,
			param.Request.Proto,
			statusColor,
			param.StatusCode,
			resetColor,
			param.Latency,
			param.ErrorMessage,
		)
	})
}

// Cors 跨域
func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		if origin != "" {
			context.Header("Access-Control-Allow-Origin", origin)
			context.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
