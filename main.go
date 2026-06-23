package main

import (
	"github.com/Sri-Varshith/cp_help_cli/cmd"
	"github.com/Sri-Varshith/cp_help_cli/internal/ai"
)

func main() {
	ai.LoadEnv()
	cmd.Execute()
}