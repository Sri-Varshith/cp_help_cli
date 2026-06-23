package cmd

import (
	"fmt"

	"github.com/Sri-Varshith/cp_help_cli/internal/parser"
	"github.com/spf13/cobra"
)

var improveCmd = &cobra.Command{
	Use:   "improve <file>",
	Short: "Suggest code improvements",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		code, err := parser.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		fmt.Printf("Successfully read %s\n", file)
		fmt.Printf("Code length: %d bytes\n", len(code))
	},
}

func init() {
	rootCmd.AddCommand(improveCmd)
}