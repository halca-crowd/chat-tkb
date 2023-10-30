package main

import (
	"encoding/json"
	"log/slog"
	"math"

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
