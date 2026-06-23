package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze <file>",
	Short: "Analyze time and space complexity",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		fmt.Printf("Analyzing %s\n", file)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
