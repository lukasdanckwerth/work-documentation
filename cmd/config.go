package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(CreateConfig)
	configCmd.AddCommand(readConfig)
}

// Config subcommand
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create / manipulate config",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Create a new config file
var CreateConfig = &cobra.Command{
	Use:   "create",
	Short: "Creates the config file",
	Run: func(cmd *cobra.Command, args []string) {

		err := viper.SafeWriteConfig()

		if err != nil {
			// error message contains enough information
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("config created")
		}
	},
}

// Read and print the config
var readConfig = &cobra.Command{
	Use:   "read",
	Short: "Read the config file",
	Run: func(cmd *cobra.Command, args []string) {

		c := viper.AllSettings()
		bs, err := yaml.Marshal(c)

		if err != nil {
			log.Fatalf("unable to marshal config to YAML: %v", err)
		}

		fmt.Println(string(bs))
	},
}
