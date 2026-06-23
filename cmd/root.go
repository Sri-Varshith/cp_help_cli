package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cphelp",
	Short: "AI-powered competitive programming assistant",
	Long:  "Analyze, debug and improve code using AI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
