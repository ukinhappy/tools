package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "日常所用的工具",
	Long: `日常所用的工具`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
