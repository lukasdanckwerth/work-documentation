package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var short bool

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&short, "short", "s", false, "Only print the version number")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + CommandName,
	Run: func(cmd *cobra.Command, args []string) {
		if short {
			fmt.Println(Version)
		} else {
			fmt.Printf("%s version %s\n", CommandName, Version)
		}
	},
}
