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
				Prompt:  "Failed to unmarshal message",
				Created: 0,
			}
		}
		res[i] = message
	}
	return res, nil
}

func chatmessageToJSON(msg []string) ([]ChatMessage, error) {
	// 文字列がMessageオブジェクトなのでJSONに変換する
	res := make([]ChatMessage, len(msg))
	for i, m := range msg {
		// Message型にUnmarshalする
		var message ChatMessage
		err := json.Unmarshal([]byte(m), &message)
		if err != nil {
			// ログに出力する
			slog.Warn("Failed to unmarshal message: %s", err)
			res[i] = ChatMessage{
				Role:    "Failed to unmarshal message",
				Content: Message{},
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

func getChatHistory(history_origin int64) ([]ChatMessage, error) {
	history := []ChatMessage{}
	// history_origin以降のチャット履歴をキャッシュから取得する
	messages, err := redis.LRange(KEY_CHAT_HISTORY, history_origin, math.MaxInt)
	if err != nil {
		return nil, err
	}
	// 文字列がMessageオブジェクトなのでJSONに変換する
	history, err = chatmessageToJSON(messages)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func saveChatHistory(history []ChatMessage) (err error) {
}

func getChatHistoryOrigin() (int64, error) {
	// キャッシュから会話ログの取り出し
	history_origin, err := redis.LRange(KEY_CHAT_HISTORY_ORIGIN, 0, 0)
	if err != nil {
		return -1, err
	}
	// 文字列がMessageオブジェクトなのでJSONに変換する
	history_origin_int, err := strconv.ParseInt(history_origin[0], 10, 64)
	if err != nil {
		return -1, err
	}
	return history_origin_int, nil
}

func setChatHistoryOrigin(history_origin int64) (err error) {
	// オブジェクトデータの組み立て
	chatObj := ContextOrigin{
		Origin: int64(history_origin),
	}

	// Message型にMarshalする
	byte_data, err := json.Marshal(chatObj)
	if err != nil {
		return err
	}

	// redisに格納するデータを作成
	insert_data := string(byte_data)
	err = redis.LPush(KEY_CHAT_HISTORY_ORIGIN, insert_data)
	if err != nil {
		return err
	}
	return nil
}

func savePresetMsg(msg ChatMessage) (err error) {
	// Message型にMarshalする
	byte_data, err := json.Marshal(msg)
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
	msgObj := ChatMessage{
		Role:    Assistant,
		Content: Message{Message: msg, Created: time.Now().Unix()},
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
func requestPrompt(msg string, history []ChatMessage) (ChatMessage, error) {

	model := "gpt-3.5-turbo-16k"
	prompt := msg

	// modelがModelのどれかに該当するか確認
	model_name, err := strToModel(model)
	if err != nil {
		return ChatMessage{}, err
	}

	openai_token := os.Getenv(API_TOKEN_ENV_NAME_OPENAI)
	openai_org_id := os.Getenv(OPENAI_ORGANIZATION_ID)
	max_tokens, err := strconv.Atoi(os.Getenv(STREAM_MAX_TOKENS_ENV_NAME))
	if err != nil {
		return ChatMessage{}, err
	}

	// OpenAIのAPIにリクエストを投げる
	log.Println("call llm_api with model: " + model + " and prompt: " + prompt)
	client := createClient(&http.Client{}, openai_token, openai_org_id)
	result := ChatMessage{}
	err = nil
	switch model_name {
	case GPT3Dot5Turbo:
		res := []ChatMessage{}
		res, err = throwChatStreamRequests(client, selectModel(model_name), prompt, max_tokens, history)
		result = res[len(res)-1]
	}
	if err != nil {
		return ChatMessage{}, err
	}

	return result, nil
}
