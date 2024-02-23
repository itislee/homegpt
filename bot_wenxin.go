package main

import (
	"fmt"
	"context"
	ernie "github.com/anhao/go-ernie"
)

type WenXinBot struct {
	client *ernie.Client
	model string
}

func NewWenXinBot(id, secret string) *WenXinBot {
	return &WenXinBot {
		client: ernie.NewDefaultClient(id, secret),
		model: "yi_34b_chat", // 默认是YI34B API免费调用
	}
}

func (b *WenXinBot)SetModel(m string) {
	b.model = m
}

func (b *WenXinBot)Name() string {
	return "文心一言"
}

func (b *WenXinBot)Ask(question string ) string {
	completion, err := b.client.CreateBaiduChatCompletion(context.Background(), ernie.BaiduChatRequest{
		Messages: []ernie.ChatCompletionMessage{
			{
				Role:    ernie.MessageRoleUser,
				Content: question,
			},
		},
		Model: b.model,
	})
	if err != nil {
		fmt.Printf("ernie bot error: %v\n", err)
		return ""
	}
	return completion.Result
}