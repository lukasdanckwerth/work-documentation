package cmd

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/lukasdanckwerth/work-documentation/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().StringVarP(&day, "day", "d", fmt.Sprint(time.Now().Day()), "Day to list")
	summaryCmd.Flags().StringVarP(&month, "month", "m", fmt.Sprint(int(time.Now().Month())), "Month to list")
	summaryCmd.Flags().StringVarP(&year, "year", "y", fmt.Sprint(time.Now().Year()), "Year to list")
}

var summaryCmd = &cobra.Command{
	Use:     "sum",
	Aliases: []string{"month"},
	Short:   "Lists to current month",
	Run: func(cmd *cobra.Command, args []string) {

		directory := model.WorkDirectory()
		monthObjs := directory.ReceiveMonth(year, month)

		var total int64
		var theMap = make(map[string]int64)

		for _, day := range monthObjs {
			for _, entry := range day.Entries {
				theMap[entry.Title] += entry.Length
				total += entry.Length
			}
		}

		keys := make([]string, 0, len(theMap))

		for key := range theMap {
			keys = append(keys, key)
		}

		sort.SliceStable(keys, func(i, j int) bool {
			return theMap[keys[i]] > theMap[keys[j]]
		})

		tw := table.NewWriter()
		tw.SetOutputMirror(os.Stdout)
		tw.AppendHeader(table.Row{"#", "task", "start", "length", "length CATS"})

		for i, title := range keys {
			tw.AppendRow(
				table.Row{
					i,
					title,
					time.UnixMilli(0),
					model.Format(theMap[title]),
					model.FormatCats(theMap[title]),
				},
			)
		}

		tw.AppendSeparator()
		tw.AppendFooter(table.Row{"", "", "Total", model.Format(total), model.FormatCats(total)})

		tw.Render()
	},
}
