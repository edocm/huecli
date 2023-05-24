package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Exists bool

func LoadConfig() {
	initConfig()
	exists, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}
	Exists = exists
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("bridge", "")
	viper.SetDefault("username", "")
	viper.SetDefault("clientkey", "")
	// TODO: only print responses if true
	viper.SetDefault("responses", true)
}

func readConfig() (bool, error) {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false, nil
		} else {
			return true, fmt.Errorf("error while reading config file: %v", err)
		}
	}
	return true, nil
}
