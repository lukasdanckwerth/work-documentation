package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CommandName = "wodo"
var Version = "0.0.2"

var rootCmd = &cobra.Command{
	Use:   fmt.Sprintf("%s [describe] [your] [task] ...", CommandName),
	Short: fmt.Sprintf("%s (work documentation) is tool to simply, fast and affortless documentate your work", CommandName),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
