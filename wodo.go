package main

import (
	"fmt"

	"github.com/lukasdanckwerth/work-documentation/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.wodo")
	viper.AddConfigPath(".")

	viper.SetDefault("user", "lukas")
	viper.SetDefault("hours", 30)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// config file not found; ignore error if desired
			// cmd.CreateConfig()
		} else {
			panic(fmt.Errorf("fatal error when reading config file: %w", err))
		}
	}

	cmd.Execute()
}
