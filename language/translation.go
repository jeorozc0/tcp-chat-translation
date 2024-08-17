package lang

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func TranslateMsg(msg, sourceLang, targetLang string) (string, error) {
	if sourceLang == targetLang {
		fmt.Printf("%s same as %s", sourceLang, targetLang)
		return msg, nil
	}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		return "", fmt.Errorf("Error while reading config file %s", err)
	}
	apiKey, ok := viper.Get("OPENAI_API_KEY").(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		return "", fmt.Errorf("Invalid type assertion")
	}
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	prompt := fmt.Sprintf("Translate the following text from %s to %s: %s", sourceLang, targetLang, msg)

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("Couldn't translate message: %v\n", err)
	}

	return resp.Choices[0].Message.Content, nil
}
