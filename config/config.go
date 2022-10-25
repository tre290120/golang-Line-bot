package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func GetLineConfig(path string) (string, string) {
	viper.SetConfigName("line_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("read line token file failed")
		os.Exit(1)
	}
	return viper.GetString("line.token"), viper.GetString("line.secret")
}
