package ginworker

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

// initConfig reads in config file and ENV variables if set.
func InitConfig() {
	Config = viper.New()
	Config.SetConfigFile("./conf/gin.yaml")
	Config.ReadInConfig()
	log.Println(Config.Get("env"))
}
