package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// calculatorCmd represents the time command
var calculatorCmd = &cobra.Command{
	Use:   "cal",
	Short: "计算器功能",
	Long:  `计算器`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// remainder represents the time command
var remainder = &cobra.Command{
	Use:   "remainder",
	Short: "余数",
	Long:  `arg1 被余数 arg2余数 `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("缺少参数")
			return
		}
		v, _ := strconv.Atoi(args[0])
		k, _ := strconv.Atoi(args[1])
		fmt.Printf("商:%d 余数:%d", v/k, v%k)
	},
}

func init() {
	rootCmd.AddCommand(calculatorCmd)
	calculatorCmd.AddCommand(remainder)
}
