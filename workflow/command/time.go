package command

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ukinhappy/go-utils/timex"
	"strconv"
)

// oa represents the time command
var timeCmd = &cobra.Command{
	Use:   "time [OPTIONS]",
	Short: "time command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var flowresult XmlItems
		var result1, result2 string
		if len(args) > 0 {
			ts, _ := strconv.Atoi(args[0])
			if ts > 0 {
				result1 = timex.Unix(int64(ts)).String()
				result2 = strconv.Itoa(ts)
			} else {
				result1 = strconv.Itoa(int(timex.String(args[0] + " " + args[1]).UnixSecond()))
				result2 = args[0] + " " + args[1]
			}
		} else {
			now := timex.Now().UnixSecond()
			result1 = timex.Unix(now).String()
			result2 = strconv.Itoa(int(timex.String(result1).UnixSecond()))
		}
		flowresult.Iterm = append(flowresult.Iterm,
			Item{
				Title:    result2,
				Subtitle: "时间点:" + result2,
				Arg:      result2,
				Valid:    "yes"},
			Item{
				Title:    result1,
				Subtitle: "时间点:" + result1,
				Arg:      result1,
				Valid:    "yes"},
		)
		b, _ := xml.Marshal(flowresult)
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
