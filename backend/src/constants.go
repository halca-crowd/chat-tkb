package main

import "fmt"

// Preset Response
var FatalErrorResponse = []byte(`{"action":"ERROR_MESSAGE","code":503,"msg":"failed to make erorr obj and occured fatal error in the server","name":"fatal error"}`)
var SystemResponse = []byte(`{"action":"SYSTEM_MESSAGE","status":"OK","error": false}`)

// 環境変数の名前
const API_PORT_VARIABLE_NAME = "LLM_API_PORT"           // APIを公開するポートを指定する環境変数の名前
const API_TOKEN_ENV_NAME_OPENAI = "OPENAI_API_TOKEN"    // OpenAIのAPIのトークンを指定する環境変数の名前
const LOG_INFO_ENV_NAME = "LOGGING_INFO"                // Infoレベルのログを出力するか否かを指定する環境変数の名前（空文字列なら出力しない）
const STREAM_MAX_TOKENS_ENV_NAME = "STREAM_MAX_TOKENS"  // デコーダモデルにEOFが出力されるまでの最大トークン数を指定する環境変数の名前
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

// Presetメッセージを格納するためのKEY
const KEY_GPT_WORD = "gpt_word"

const CONNECTION_PATH = "connection_path"

const ACTION_CHAT_MESSAGE = "chat_message"

const RES_GPT_MESSAGE = "gpt_message"
const TEST_CHAT_MESSAGE = "test_chat_message"

const ACTION_ERROR = "ERROR_MESSAGE"
