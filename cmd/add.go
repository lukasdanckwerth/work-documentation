package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lukasdanckwerth/work-documentation/model"
	"github.com/spf13/cobra"
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
	Use:   "add <describe> <your> <task>...",
	Short: "Add a task done today",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Not taksk provided.")
			os.Exit(0)
		}

		var title = strings.Join(args, " ")
		var entry = model.Entry{
			Title:   title,
			Created: time.Now().UnixMilli(),
		}

		if length != "" {
			var length, error = model.ParseTime(length)
			if error != nil {
				fmt.Println(error)
				os.Exit(1)
			}

			entry.Length = length
		}

		directory := model.WorkDirectory()
		dayObj := directory.ReceiveDay(year, month, day)
		dayObj.AddEntry(entry)

		directory.WriteWorkday(dayObj, year, month, day)
	},
}
