package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

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
func throwChatStreamRequests(client *openai.Client, model string, prompt string, max_tokens int, history []openai.ChatCompletionMessage) ([]openai.ChatCompletionMessage, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
	// if the history exists, insert history at first index of messages
	if len(history) > 0 {
		messages = append(messages, history...)
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
		return messages, err
	}

	defer stream.Close()
	response_text := ""
	for {
		resp_i, err := stream.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return messages, err
			}
			break
		}
		fmt.Println(resp_i)
		response_text += resp_i.Choices[0].Delta.Content
	}
	response := append(
		messages,
		openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: response_text,
		})

	return response, nil
}

// モデル選択
func selectModel(model_name ModelName) string {
	switch model_name {
	case GPT3Dot5Turbo:
		return openai.GPT3Dot5Turbo
	case Davinci:
		return openai.GPT3Davinci
	default:
		return ""
	}
}
