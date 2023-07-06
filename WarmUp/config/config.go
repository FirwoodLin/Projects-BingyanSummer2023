package config

import (
	"fmt"

	"github.com/spf13/viper"
	// "WarmUp/logger"
)

func ReadIn() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("[fatal error] config file: %s \n", err)
		// logger.
		// logger.lger.Fatal("fatal error config file: %s \n", err)
	}
}
