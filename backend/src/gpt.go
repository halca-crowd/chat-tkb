package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"notchman.tech/chat-tkb/src/redis"
)

func messagesToJSON(msg []string) ([]Message, error) {
	// 文字列がMessageオブジェクトなのでJSONに変換する
	res := make([]Message, len(msg))
	for i, m := range msg {
		// Message型にUnmarshalする
		var message Message
		err := json.Unmarshal([]byte(m), &message)
		if err != nil {
			// ログに出力する
			slog.Warn("Failed to unmarshal message: %s", err)
			res[i] = Message{
				Message: "Failed to unmarshal message",
				Prompt: "Failed to unmarshal message",
				Created: 0,
			}
		}
		res[i] = message
	}
	return res, nil
}

func buildPresetUserMessages() ([]Message, error) {
	//キャッシュの作成
	messages, err := redis.LRange(KEY_GPT_WORD, 0, math.MaxInt)
	if err != nil {
		return nil, err
	}

	presetMessages := messages
	// 文字列がMessageオブジェクトなのでJSONに変換する
	return messagesToJSON(presetMessages)
}

func savePresetMsg(prompt string,msg string) (err error) {
	// オブジェクトデータの組み立て
	msgObj := Message{
		Message: msg,
		Prompt: prompt,
		Created: time.Now().Unix(),
	}

	// Message型にMarshalする
	byte_data, err := json.Marshal(msgObj)
	if err != nil {
		return err
	}

	// redisに格納するデータを作成
	insert_data := string(byte_data)
	err = redis.LPush(KEY_GPT_WORD, insert_data)
	if err != nil {
		return err
	}
	return nil
}

func savePromptMsg(msg string) (err error) {
	// オブジェクトデータの組み立て
	msgObj := PromptData{
		Prompt: msg,
		Created: time.Now().Unix(),
	}

	// Message型にMarshalする
	byte_data, err := json.Marshal(msgObj)
	if err != nil {
		return err
	}

	// redisに格納するデータを作成
	insert_data := string(byte_data)
	err = redis.RPush(KEY_GPT_WORD, insert_data)
	if err != nil {
		return err
	}
	return nil
}

// LLM APIのコントローラ
func requestPrompt(msg string) (string, error) {

	model := "gpt3.5-turbo"
	prompt := msg

	// modelがModelのどれかに該当するか確認
	model_name, err := strToModel(model)
	if err != nil {
		return "", err
	}

	openai_token := os.Getenv(API_TOKEN_ENV_NAME_OPENAI)
	openai_org_id := os.Getenv(OPENAI_ORGANIZATION_ID)
	max_tokens, err := strconv.Atoi(os.Getenv(STREAM_MAX_TOKENS_ENV_NAME))
	if err != nil {
		return "", err
	}

	// OpenAIのAPIにリクエストを投げる
	log.Println("call llm_api with model: " + model + " and prompt: " + prompt)
	client := createClient(&http.Client{}, openai_token, openai_org_id)
	result := ""
	err = nil
	switch model_name {
	case Davinci:
		result, err = throwStreamRequest(client, selectModel(model_name), prompt, max_tokens)
	case GPT3Dot5Turbo:
		result, err = throwChatStreamRequest(client, selectModel(model_name), prompt, max_tokens)
	}
	if err != nil {
		return "", err
	}

	return result, nil
}
