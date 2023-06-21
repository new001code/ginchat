package common

import (
	"fmt"
	"ginchat/util"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Log = &util.MyLog{}

func init() {
	Log.Debugln("start load configuration file")
	viper.SetConfigFile("./config/application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading configuration file: %s", err))
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		Log.Debugf("config file changed: %s \n", e.Name)
	})
	Log.Debugln("success load configuration file")
}
