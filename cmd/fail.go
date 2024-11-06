package cmd

import (
	"github.com/gkwa/britishinput/core"
	"github.com/spf13/cobra"
)

var failCmd = &cobra.Command{
	Use:   "fail",
	Short: "Always exits with failure",
	Long: `The fail command will always exit with failure (status code 1).
This can be useful for testing error handling and failure scenarios.`,
	Run: core.RunFailer,
}

func init() {
	rootCmd.AddCommand(failCmd)
}
