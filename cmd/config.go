package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	},
}

var readConfig = &cobra.Command{
	Use:   "write",
	Short: "Read the config file",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.SafeWriteConfig()

		if err != nil {
			panic(err)
		}
	},
}
