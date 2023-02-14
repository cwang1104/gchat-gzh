package cgtp

import (
	"context"
	"fmt"
	"gchat-gzh/pkg/logger"
	gogpt "github.com/sashabaranov/go-gpt3"
	"time"
)

const (
	apikey   = ""
	maxToken = 500
	timeout  = time.Second * 60 * 30
)

type ChatController struct {
	maxToken int
}

func NewChatGtp() *ChatController {
	return &ChatController{
		maxToken: maxToken,
	}
}

func (c *ChatController) SetMaxToken(maxToken int) {
	c.maxToken = maxToken
}

func (c *ChatController) GetMaxToken() int {
	return c.maxToken
}

func (c *ChatController) GetResponse(prompt string) (resp string, err error) {
	fmt.Println("收到问题：", prompt)
	client := gogpt.NewClient(apikey)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        500,
		Prompt:           prompt,
		Temperature:      0.9,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
	}
	response, err := client.CreateCompletion(ctx, req)
	if err != nil {
		logger.Log.Error("request error", err)
		return "", err
	}
	fmt.Printf("接收到问题：%s,\n答案：%s\n", prompt, response.Choices[0].Text)
	return response.Choices[0].Text, nil
}
