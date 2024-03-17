// Copyright 2024 Lukas Danckwerth
//

package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lukasdanckwerth/work-documentation/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var year string
var month string
var day string
var length string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&length, "length", "l", "", "Length of the task")
	addCmd.Flags().StringVarP(&day, "day", "d", fmt.Sprint(time.Now().Day()), "Day of the task")
	addCmd.Flags().StringVarP(&month, "month", "m", fmt.Sprint(int(time.Now().Month())), "Month of the task")
	addCmd.Flags().StringVarP(&year, "year", "y", fmt.Sprint(time.Now().Year()), "Year of the task")
}

var addCmd = &cobra.Command{
	Use:     "add <describe> <your> <task>...",
	Short:   "Add a task done today",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {

		var title = strings.Join(args, " ")
		if len(args) == 0 {
			title = promtTaskTitle()
		}

		if title == "" {
			fmt.Println("no task title provided")
			os.Exit(1)
		}

		var entry = model.Entry{
			Title:   title,
			Created: time.Now().UnixMilli(),
			UUID:    uuid.NewString(),
		}

		if length == "" {
			length = promptLength()
		}

		var length, error = model.ParseTime(length)
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}

		entry.Length = length

		d := model.WorkDirectory()
		wd := d.ReceiveWorkday(year, month, day)
		wd.AddEntry(entry)

		d.WriteWorkday(wd, year, month, day)
	},
}

func promtTaskTitle() string {

	var taskOptions = viper.GetStringSlice("tasks")

	prompt := promptui.Select{
		Label: "Select task",
		Items: taskOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func promptLength() string {

	var lengthOptions = viper.GetStringSlice("lengths")

	prompt := promptui.Select{
		Label: "Select length",
		Items: lengthOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
