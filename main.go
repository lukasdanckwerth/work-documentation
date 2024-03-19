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

	// viper.SetDefault("work-hours", 30)
	// viper.SetDefault("work-day-hours", 6)
	viper.SetDefault("tasks", []string{"task 1", "task 2"})
	viper.SetDefault("lengths", []string{"15m", "30m", "45m", "1h", "1:30h", "2h", "2:30h", "3h",
		"4h", "5h", "6h", "7h", "8h"})

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cmd.CreateConfig.Run(cmd.CreateConfig, []string{})
		} else {
			panic(fmt.Errorf("fatal error when reading config file: %w", err))
		}
	}

	cmd.Execute()
}
