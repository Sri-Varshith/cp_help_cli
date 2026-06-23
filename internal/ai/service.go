package ai

import (
	"github.com/Sri-Varshith/cp_help_cli/internal/parser"
)

func ProcessFile(
	file string,
	builder func(string) string,
) (string, error) {

	code, err := parser.ReadFile(file)
	if err != nil {
		return "", err
	}

	prompt := builder(code)

	return Ask(prompt)
}
