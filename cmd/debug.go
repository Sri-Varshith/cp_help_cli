package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug <file>",
	Short: "Find bugs and issues in code",
	Long:  "Analyze a source file and identify bugs, runtime errors, edge cases, and logic issues.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		fmt.Printf("Debugging %s\n", file)

		// TODO:
		// 1. Read file contents
		// 2. Send code to AI
		// 3. Display bug report
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}