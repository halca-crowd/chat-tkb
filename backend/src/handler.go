package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"time"

	"github.com/pkg/errors"
	"notchman.tech/chat-tkb/src/redis"
)

func count(target string) (value int, err error) {
	err = redis.AddValue(target)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to add connection")
	}

	value, err = redis.DeclValue(target)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to decl connection")
	}
	return
}

func handler(s []byte) []byte {
	var requestObject Request
	if err := json.Unmarshal(s, &requestObject); err != nil {
		log.Println(err)
		return errorResponseFactory("faile to parse json", 503, err.Error())

	}

	// 各アクションケースに応じて処理を行う
	switch {
	case requestObject.Action == TEST_CHAT_MESSAGE:
		r := []byte(`{"action":"ACTION_RECV_GPT","user_id":"examper-user-id","message":"hogehoge"}`)
		return r

	case requestObject.Action == ACTION_CHAT_MESSAGE:
		// LLM APIにリクエストを送信する
		res, err := requestPrompt(requestObject.Message)

		if err != nil {
			slog.Error(err.Error())
			_ = savePresetMsg(requestObject.Message,"failed to fetch openai api")
			return messageResponseFactory(requestObject.Message,"failed to fetch openai api")
			// return errorResponseFactory("faile to send message", 503, "data is not json object")
		}
		err = savePresetMsg(requestObject.Message,res)
		if err != nil{
			slog.Info("failed to save preset message")
		}
		err = savePromptMsg(requestObject.Message)
		if err != nil{
			slog.Info("failed to save preset message")
		}
		return messageResponseFactory(requestObject.Message,res)
	case requestObject.Action == ACTION_FORCE_RESET:
		// 強制削除メッセージを送信
		return forceResetMessageFactory()
	default:
		return errorResponseFactory("faile to execute action", 503, "no such action type")
	}

}

func errorResponseFactory(name string, code int, msg string) []byte {
	errRes := ErrorObject{
		Action: ACTION_ERROR,
		Name:   name,
		Code:   code,
		Msg:    msg,
	}
	res, err := json.Marshal(errRes)
	if err != nil {
		return FatalErrorResponse
	}
	return res
}

func forceResetMessageFactory()[]byte{
	forceResetMessage := ForceResetMessage{
		Action: ACTION_FORCE_RESET,
	}
	res, err := json.Marshal(forceResetMessage)
	if err != nil {
		return FatalErrorResponse
	}
	return res
}

func messageResponseFactory(inputMsg string,outputMsg string) []byte {
	// 現在時刻を取得
	current_time := time.Now().Unix()
	resObj := ChatResponse{
		Action:  RES_GPT_MESSAGE,
		Message: outputMsg,
		Created: current_time,
		Prompt: inputMsg,
	}
	res, err := json.Marshal(resObj)
	if err != nil {
		return FatalErrorResponse
	}
	return res
}
