package main

import "fmt"

// 環境変数の名前
const API_PORT_VARIABLE_NAME = "LLM_API_PORT"          // APIを公開するポートを指定する環境変数の名前
const API_TOKEN_ENV_NAME_OPENAI = "OPENAI_API_TOKEN"   // OpenAIのAPIのトークンを指定する環境変数の名前
const LOG_INFO_ENV_NAME = "LOGGING_INFO"               // Infoレベルのログを出力するか否かを指定する環境変数の名前（空文字列なら出力しない）
const STREAM_MAX_TOKENS_ENV_NAME = "STREAM_MAX_TOKENS" // デコーダモデルにEOFが出力されるまでの最大トークン数を指定する環境変数の名前
const OPENAI_ORGANIZATION_ID = "OPENAI_ORGANIZATION_ID" // OpenAIのOrganization IDを指定する環境変数の名前

// エンドポイント名
const ENDPOINT_LLM_API = "/llm_api"

// リクエストのパラメータ名
const MODEL_PARAM_NAME = "model"   // モデル名
const PROMPT_PARAM_NAME = "prompt" // プロンプト

// モデル名
type ModelName string // モデル名のEnum

const (
	GPT3Dot5Turbo ModelName = "gpt3.5-turbo" // GPT3.5 Turbo
	Davinci       ModelName = "davinci"      // Davinci
)

func strToModel(model string) (ModelName, error) {
	switch model {
	case string(GPT3Dot5Turbo):
		return GPT3Dot5Turbo, nil
	case string(Davinci):
		return Davinci, nil
	default:
		return "", fmt.Errorf("Invalid model name: %s", model)
	}
}

type ErrorLevel int // ログのレベルのEnum

const (
	Info    ErrorLevel = iota // Info（デフォルトでは出力しない）
	Debug                     // Debug
	Warning                   // Warning
	Error                     // Error（ログ出力後に終了）
)

// ErrorLevelを対応する文字列に変換する
func (el ErrorLevel) String() string {
	switch el {
	case Info:
		return "Info"
	case Debug:
		return "Debug"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	default:
		return ""
	}
}