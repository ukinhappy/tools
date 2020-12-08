package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version info",
	Long: `version info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
