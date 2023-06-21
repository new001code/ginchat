package common

import (
	"fmt"
	"ginchat/util"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	util.DebugLogger.Println("start load configuration file")
	viper.SetConfigFile("./config/application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading configuration file: %s", err))
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		util.DebugLogger.Printf("config file changed: %s \n", e.Name)
	})
	util.DebugLogger.Println("success load configuration file")
}
