package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIBot struct {
	client *openai.Client
	model  string
}

func NewOpenAIBot(secret string) *OpenAIBot {
	config := openai.DefaultConfig(secret)
	config.BaseURL = "https://aihubmix.com/v1"
	return &OpenAIBot{
		client: openai.NewClientWithConfig(config),
		model:  openai.GPT3Dot5Turbo,
	}
}

func (b *OpenAIBot) SetModel(m string) {
	b.model = m
}

func (b *OpenAIBot) Name() string {
	return "GPT"
}

func (b *OpenAIBot) Ask(question string) string {
	resp, err := b.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: b.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	},
	)
	if err != nil {
		fmt.Printf("openai bot error: %v\n", err)
		return ""
	}
	return resp.Choices[0].Message.Content
}
