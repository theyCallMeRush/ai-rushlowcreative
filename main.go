package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: " business name ideas for a painting",
				},
			},
		},
	)

	if err != nil {
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
