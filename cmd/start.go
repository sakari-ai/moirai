package cmd

import (
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start given service",
	Long:  "Start given service",
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func Register(cmd *cobra.Command) {
	startCmd.AddCommand(cmd)
}
