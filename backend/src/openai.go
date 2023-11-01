package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// OpenAIのclientを作成する
func createClient(http_client *http.Client, token string, orgID string) *openai.Client {
	config := openai.DefaultConfig(token)
	if orgID != "" {
		config.OrgID = orgID
	}
	config.HTTPClient = http_client
	client := openai.NewClientWithConfig(config)

	return client
}

// デコーダモデルにリクエストを投げる
func throwRequest(client *openai.Client, model string, prompt string) (openai.CompletionResponse, error) {
	return client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  model,
			Prompt: prompt,
		},
	)
}

// デコーダモデルにEOFが出力されるまでリクエストを投げる
func throwStreamRequest(client *openai.Client, model string, prompt string, max_tokens int) (string, error) {
	stream, err := client.CreateCompletionStream(
		context.Background(),
		openai.CompletionRequest{
			Model:     model,
			Prompt:    prompt,
			Stream:    true,
			MaxTokens: max_tokens,
		},
	)
	if err != nil {
		return "", err
	}

	defer stream.Close()
	response := ""
	for {
		resp_i, err := stream.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return "", err
			}
			break
		}
		response += resp_i.Choices[0].Text
	}
	return response, nil
}

// 対話形式のモデルにリクエストを投げる
func throwChatRequest(client *openai.Client, model string, prompt string, chatLog []string) (openai.ChatCompletionResponse, error) {
	return client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
}

// 対話形式のモデルにEOFが出力されるまでリクエストを投げる
func throwChatStreamRequest(client *openai.Client, model string, prompt string, max_tokens int) (string, error) {
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Stream:    true,
			MaxTokens: max_tokens,
		},
	)
	if err != nil {
		return "", err
	}

	defer stream.Close()
	response := ""
	for {
		resp_i, err := stream.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return "", err
			}
			break
		}
		fmt.Println(resp_i)
		response += resp_i.Choices[0].Delta.Content
	}
	return response, nil
}

// 対話形式のモデルにEOFが出力されるまでリクエストを投げる。複数回の対話に対応
func throwChatStreamRequests(client *openai.Client, model string, prompt string, max_tokens int, history []ChatMessage) ([]ChatMessage, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
	// if the history exists, insert history at first index of messages
	if len(history) > 0 {
		for _, h := range history {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    roleToOpenAIRole(h.Role),
				Content: h.Content.Message,
			})
		}
	}
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     model,
			Messages:  messages,
			Stream:    true,
			MaxTokens: max_tokens,
		},
	)
	if err != nil {
		return history, err
	}

	defer stream.Close()
	response_text := ""
	for {
		resp_i, err := stream.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return history, err
			}
			break
		}
		fmt.Println(resp_i)
		response_text += resp_i.Choices[0].Delta.Content
	}
	response := append(
		history,
		ChatMessage{
			Role: openai.ChatMessageRoleAssistant,
			Content: Message{
				Message: response_text,
				Created: time.Now().Unix(),
			},
		},
	)

	return response, nil
}

// モデル選択
func selectModel(model_name ModelName) string {
	switch model_name {
	case GPT3Dot5Turbo:
		return openai.GPT3Dot5Turbo
	case GPT3Dot5Turbo16k:
		return openai.GPT3Dot5Turbo16K
	case Davinci:
		return openai.GPT3Davinci
	default:
		return ""
	}
}

// ロール変換 to OpenAI
func roleToOpenAIRole(role Role) string {
	switch role {
	case System:
		return openai.ChatMessageRoleSystem
	case User:
		return openai.ChatMessageRoleUser
	case Assistant:
		return openai.ChatMessageRoleAssistant
	default:
		return openai.ChatMessageRoleUser
	}
}
