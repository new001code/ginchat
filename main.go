package main

import (
	"fmt"
	_ "ginchat/common"
	"ginchat/router"
	_ "ginchat/util"

	"github.com/spf13/viper"
)

func main() {
	r := router.Router()
	u := fmt.Sprintf(":%d", viper.GetInt("port"))
	r.Run(u)
}
