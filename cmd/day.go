package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lukasdanckwerth/work-documentation/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayCmd)
	dayCmd.Flags().StringVarP(&day, "day", "d", fmt.Sprint(time.Now().Day()), "Day to list")
	dayCmd.Flags().StringVarP(&month, "month", "m", fmt.Sprint(int(time.Now().Month())), "Month to list")
	dayCmd.Flags().StringVarP(&year, "year", "y", fmt.Sprint(time.Now().Year()), "Year to list")
}

var dayCmd = &cobra.Command{
	Use:     "day",
	Aliases: []string{"d", "l", "ls"},
	Short:   "Lists to current day",
	Run: func(cmd *cobra.Command, args []string) {

		directory := model.WorkDirectory()
		dayObj := directory.ReceiveWorkday(year, month, day)

		tw := table.NewWriter()
		tw.SetOutputMirror(os.Stdout)
		tw.AppendHeader(table.Row{"#", "task", "start", "length"})

		var total int64

		for i, e := range dayObj.Entries {
			tw.AppendRow(table.Row{
				i + 1,
				e.Title,
				e.CreatedTimeFormatted(),
				e.LenghtFormatted(),
			})
			total += e.Length
		}

		tw.AppendSeparator()
		tw.AppendFooter(table.Row{"", "Total", "", model.Format(total)})

		tw.Render()
	},
}
