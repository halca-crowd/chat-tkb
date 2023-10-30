package main

import (
	"encoding/json"
	"log"
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
			log.Println(err)
			return errorResponseFactory("faile to send message", 503, "data is not json object")

		}
		return messageResponseFactory(res)

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

func messageResponseFactory(msg string) []byte {
	// 現在時刻を取得
	current_time := time.Now().Unix()
	resObj := ChatResponse{
		Action:  RES_GPT_MESSAGE,
		Message: msg,
		Created: current_time,
	}
	res, err := json.Marshal(resObj)
	if err != nil {
		return FatalErrorResponse
	}
	return res
}
