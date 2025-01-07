// Package configs implements the functions, types, and interfaces for the module.
package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var ConfigFile string

func InitConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(wd)
		viper.SetConfigType("toml")
		viper.SetConfigName(".markdown.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
