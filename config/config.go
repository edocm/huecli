package config

import "github.com/spf13/viper"

var Exists bool

func LoadConfig() {
	initConfig()
	Exists = readConfig()
}

func initConfig() {
	viper.SetConfigFile("./config.yaml")
	viper.SetDefault("username", "")
	viper.SetDefault("clientkey", "")
}

func readConfig() bool {
	if err := viper.ReadInConfig(); err != nil {
		return false
	}
	return true
}
