package core

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Components map[string]ComponentConfig
}

type ComponentConfig struct {
	Command []string
}

// Version of application
var Version = "development"

// Config of application runtime
var Config AppConfig

func ReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("fatal error bad config file: %s", err))
	}
}
