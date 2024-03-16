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
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&day, "day", "d", fmt.Sprint(time.Now().Day()), "Day to list")
	listCmd.Flags().StringVarP(&month, "month", "m", fmt.Sprint(int(time.Now().Month())), "Month to list")
	listCmd.Flags().StringVarP(&year, "year", "y", fmt.Sprint(time.Now().Year()), "Year to list")
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"day"},
	Short:   "List to current day",
	Run: func(cmd *cobra.Command, args []string) {

		directory := model.WorkDirectory()
		dayObj := directory.ReceiveDay(year, month, day)

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "task", "start", "length"})

		for i, v := range dayObj.Entries {
			t.AppendRow(table.Row{i + 1, v.Title, time.UnixMilli(v.Start), v.LenghtFormatted()})
		}

		// t.AppendRows([]table.Row{
		// 	{1, "Arya", "Stark", 3000},
		// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
		// })

		// t.AppendSeparator()
		// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
		// t.AppendFooter(table.Row{"", "", "Total", 10000})

		t.Render()
	},
}
