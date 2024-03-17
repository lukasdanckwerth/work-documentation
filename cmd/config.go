package cmd

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(createConfig)
	configCmd.AddCommand(readConfig)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create / manipulate config",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createConfig = &cobra.Command{
	Use:   "create",
	Short: "Creates the config file",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Prompt{
			Label: "Do you want to create a config file?",
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		err = viper.SafeWriteConfig()
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
