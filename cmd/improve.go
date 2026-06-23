package cmd

import (
	"fmt"

	"github.com/Sri-Varshith/cp_help_cli/internal/ai"
	"github.com/Sri-Varshith/cp_help_cli/internal/parser"
	"github.com/Sri-Varshith/cp_help_cli/internal/output"

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
			fmt.Println(err)
			return
		}

		prompt := ai.BuildImprovePrompt(code)

		response, err := ai.Ask(prompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		output.PrintResponse(response)
	},
}

func init() {
	rootCmd.AddCommand(improveCmd)
}