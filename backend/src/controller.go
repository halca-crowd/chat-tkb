package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// リクエストボディの構造体
type RequestBody struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// indexのコントローラ(ヘルスチェック用)
func index(writer http.ResponseWriter, request *http.Request) {
	logger := NewHTTPLogger(os.Getenv(LOG_INFO_ENV_NAME) != "")

	if request.URL.Path != "/" {
		status := 404
		writer.WriteHeader(status)
		err := fmt.Errorf("no such endpoint: %s", request.URL.Path)
		logger.LoggingHTTPError(status, err)
		return
	}

	if request.Method != "GET" {
		status := 405
		writer.WriteHeader(status)
		err := fmt.Errorf("this endpoint allows only GET method but recieve %s", request.Method)
		logger.LoggingHTTPError(status, err)
		return
	}
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.Write([]byte("Hello, FusionComp LLM-API!"))
}

func preset(writer http.ResponseWriter, request *http.Request) {
	res, err := buildPresetUserMessages()
	if err != nil {
		status := 500
		writer.WriteHeader(status)
		return
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(res)
}

// LLM APIのコントローラ
func llm_api(writer http.ResponseWriter, request *http.Request) {
	logger := NewHTTPLogger(os.Getenv(LOG_INFO_ENV_NAME) != "")

	if request.Method != "GET" {
		status := 405
		writer.WriteHeader(status)
		err := fmt.Errorf("this endpoint allows only GET method but recieve %s", request.Method)
		logger.LoggingHTTPError(status, err)
		return
	}

	model := ""
	prompt := ""

	// Get Request Parameter from Body (json) when Content-Type is application/json
	if request.Header.Get("Content-Type") == "application/json" {
		decoder := json.NewDecoder(request.Body)
		var rb RequestBody
		err := decoder.Decode(&rb)
		if err != nil {
			status := 400
			writer.WriteHeader(status)
			logger.LoggingHTTPError(status, err)
			return
		}
		model = rb.Model
		prompt = rb.Prompt
	} else {
		request.ParseForm()
		model = strings.Join(request.Form[MODEL_PARAM_NAME], "")
		if model == "" {
			status := 400
			err := fmt.Errorf("you should set the parameter %s", MODEL_PARAM_NAME)
			writer.WriteHeader(status)
			logger.LoggingHTTPError(status, err)
			return
		}

		prompt = strings.Join(request.Form[PROMPT_PARAM_NAME], "")
		if prompt == "" {
			status := 400
			err := fmt.Errorf("you should set the parameter %s", PROMPT_PARAM_NAME)
			writer.WriteHeader(status)
			logger.LoggingHTTPError(status, err)
			return
		}
	}
	// modelがModelのどれかに該当するか確認
	model_name, err := strToModel(model)
	if err != nil {
		status := 400
		writer.WriteHeader(status)
		logger.LoggingHTTPError(status, err)
		return
	}

	openai_token := os.Getenv(API_TOKEN_ENV_NAME_OPENAI)
	openai_org_id := os.Getenv(OPENAI_ORGANIZATION_ID)
	max_tokens, err := strconv.Atoi(os.Getenv(STREAM_MAX_TOKENS_ENV_NAME))
	if err != nil {
		status := 500
		writer.WriteHeader(status)
		logger.LoggingHTTPError(status, err)
		return
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
		status := 500
		writer.WriteHeader(status)
		logger.LoggingHTTPError(status, err)
		return
	}

	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.Write([]byte(result))
}
