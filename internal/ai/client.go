package ai

import (
	"context"
	"errors"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
)

func Ask(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY not found")
	}

	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-5"
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	resp, err := client.Responses.New(
		context.Background(),
		responses.ResponseNewParams{
			Model: model,
			Input: responses.ResponseNewParamsInputUnion{
				OfString: openai.String(prompt),
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.OutputText(), nil
}
