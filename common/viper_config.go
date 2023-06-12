package common

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	log.Println("start load configuration file")
	viper.SetConfigFile("./config/application.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading configuration file: %s \n", err))
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed: %s \n", e.Name)
	})
	log.Println("success load configuration file")
}
