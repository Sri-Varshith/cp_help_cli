package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var improveCmd = &cobra.Command{
	Use:   "improve <file>",
	Short: "Suggest code improvements",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		fmt.Printf("Improving %s\n", file)
	},
}

func init() {
	rootCmd.AddCommand(improveCmd)
}
