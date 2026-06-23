package cmd

import (
	"fmt"

	"github.com/Sri-Varshith/cp_help_cli/internal/ai"
	"github.com/Sri-Varshith/cp_help_cli/internal/parser"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze <file>",
	Short: "Analyze time and space complexity",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		code, err := parser.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		prompt := ai.BuildAnalyzePrompt(code)

		response, err := ai.Ask(prompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
