package cmd

import (
	"fmt"
	"time"

	"github.com/lukasdanckwerth/work-documentation/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dummyDataCmd)
}

func createEntry(title string, length int64) *model.Entry {
	return &model.Entry{
		Title:  title,
		Length: length,
	}
}

var dummyDataCmd = &cobra.Command{
	Use:   "dummy",
	Short: "Creates some dummy data",
	Run: func(cmd *cobra.Command, args []string) {

		dYear := fmt.Sprint(time.Now().Year())
		dMonth := fmt.Sprint(int(time.Now().Month()))
		dDay := fmt.Sprint(time.Now().Day())
		dYesterday := fmt.Sprint(time.Now().Day() - 1)

		directory := model.WorkDirectory()

		dayObj := directory.ReceiveDay(dYear, dMonth, dDay)
		dayObj.AddEntry(*createEntry("fiss amt 33", 240))
		dayObj.AddEntry(*createEntry("jf", 120))
		directory.WriteWorkday(dayObj, dYear, dMonth, dDay)

		yesterdayObj := directory.ReceiveDay(dYear, dMonth, dYesterday)
		yesterdayObj.AddEntry(*createEntry("fiss amt 13", 200))
		yesterdayObj.AddEntry(*createEntry("veranstaltungskalender", 200))
		directory.WriteWorkday(yesterdayObj, dYear, dMonth, dYesterday)

	},
}
