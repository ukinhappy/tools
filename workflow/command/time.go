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
		var result string

		if len(args) > 0 {
			ts, _ := strconv.Atoi(args[0])
			if ts > 0 {
				result = timex.Unix(int64(ts)).String()
			} else {
				result = strconv.Itoa(int(timex.String(args[0] + " " + args[1]).UnixSecond()))
			}

		} else {
			result = strconv.Itoa(int(timex.Now().UnixSecond()))
		}
		flowresult.Iterm = append(flowresult.Iterm,
			Item{
				Title:    result,
				Subtitle: "时间点:" + result,
				Arg:      result,
				Valid:    "yes"})
		b, _ := xml.Marshal(flowresult)
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
}
