package cmd

import (
	"github.com/gkwa/britishinput/core"
	"github.com/spf13/cobra"
)

var passCmd = &cobra.Command{
	Use:   "pass",
	Short: "Always exits with success",
	Long: `The pass command will always exit with success (status code 0). 
This can be useful for testing success scenarios and verification.`,
	Run: core.RunPasser,
}

func init() {
	rootCmd.AddCommand(passCmd)
}
