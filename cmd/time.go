/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ukinhappy/go-utils/timex"
	"strconv"
	"strings"
	"time"
)

// timeCmd represents the time command
var nowTimeCmd = &cobra.Command{
	Use:   "time [OPTIONS] [COMMANDS]",
	Short: "timestamp date tools",
	Long:  `timestamp date tools include now,begindate,enddate,beginweek,endweek`,
	Run: func(cmd *cobra.Command, args []string) {
		//默认
		if len(args) > 0 {
			arg := strings.Join(args, " ")
			tm, _ := strconv.Atoi(arg)
			if tm > 0 {
				fmt.Println(time.Unix(int64(tm), 0).Format("2006-01-02 15:04:05"))
			} else {
				t, _ := time.Parse("2006-01-02 15:04:05", arg)
				fmt.Println(t.Unix())
			}
		} else {
			fmt.Println(fmt.Sprintf("[%d] : [%s]", time.Now().Unix(), time.Now().Format("2006-01-02 15:04:05")))
		}

	},
}

// timeCmd represents the time command
var beginDateTimeCmd = &cobra.Command{
	Use:   "begindate",
	Short: "begin date time",
	Long:  `begin date time`,
	Run: func(cmd *cobra.Command, args []string) {
		beginDateTm := timex.Now().BeginOfDate()
		fmt.Println(fmt.Sprintf("[%d] : [%s]", beginDateTm.T.Unix(), beginDateTm.T.Format("2006-01-02 15:04:05")))
	},
}

// timeCmd represents the time command
var endDateTimeCmd = &cobra.Command{
	Use:   "enddate",
	Short: "end date time",
	Long:  `end date time`,
	Run: func(cmd *cobra.Command, args []string) {
		endDateTime := timex.Now().EndOfDate()
		fmt.Println(fmt.Sprintf("[%d] : [%s]", endDateTime.T.Unix(), endDateTime.T.Format("2006-01-02 15:04:05")))
	},
}

// timeCmd represents the time command
var beginWeekTimeCmd = &cobra.Command{
	Use:   "beginweek",
	Short: "begin week time",
	Long:  `begin week time`,
	Run: func(cmd *cobra.Command, args []string) {
		beginWeekTm := timex.Now().BeginOfWeek()
		fmt.Println(fmt.Sprintf("[%d] : [%s]", beginWeekTm.T.Unix(), beginWeekTm.T.Format("2006-01-02 15:04:05")))
	},
}

// timeCmd represents the time command
var endWeekTimeCmd = &cobra.Command{
	Use:   "endweek",
	Short: "end week time",
	Long:  `end week time`,
	Run: func(cmd *cobra.Command, args []string) {
		endWeekTime := timex.Now().EndOfWeek()
		fmt.Println(fmt.Sprintf("[%d] : [%s]", endWeekTime.T.Unix(), endWeekTime.T.Format("2006-01-02 15:04:05")))
	},
}

func init() {
	rootCmd.AddCommand(nowTimeCmd)
	nowTimeCmd.AddCommand(beginDateTimeCmd)
	nowTimeCmd.AddCommand(endDateTimeCmd)
	nowTimeCmd.AddCommand(beginWeekTimeCmd)
	nowTimeCmd.AddCommand(endWeekTimeCmd)

}
