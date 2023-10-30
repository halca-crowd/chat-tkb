package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"notchman.tech/chat-tkb/src/cached"
)

// 起動モードの定数を管理
const (
	DEVELOPMENT = "dev"
	PRODUCTION  = "prod"
)

// APIサーバのメイン
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()
	//キャッシュの作成と導通テスト（他のやり方があれば直す）
	cacheManager := cached.NewMemcached("memcached:11211")
	//1分だけ有効なテスト値を挿入する
	cacheManager.SaveFor(time.Minute, "test", []byte("test"))
	testValue, err := cacheManager.Get("test")
	if len(testValue) == 0 || err != nil {
		log.Fatalln("failed to connect memcached")
	}
	//ローカルなどの開発環境ではNewRelicのエージェントを作成しない
	// if mode == PRODUCTION {
	// 	//newrelicのライセンスキーを取得
	// 	newrelicLicenceKey := os.Getenv("NEWRELIC_LICENCE_KEY")
	// 	if len(newrelicLicenceKey) == 0 {
	// 		logger.Error("failed to load newrelic env")
	// 	}
	// 	//newrelicエージェントの作成
	// 	app, err := newrelic.NewApplication(
	// 		newrelic.ConfigAppName("fusioncomp-llm-api"),
	// 		newrelic.ConfigLicense(newrelicLicenceKey),
	// 		newrelic.ConfigAppLogForwardingEnabled(true),
	// 	)
	// 	if err != nil {
	// 		logger.Error("failed to create newrelic agent")
	// 	}

	// 	//newrelicエージェントを導入したhttpハンドラを設定
	// 	mux.HandleFunc(newrelic.WrapHandleFunc(app, ENDPOINT_LLM_API, llm_api))
	// 	mux.HandleFunc(newrelic.WrapHandleFunc(app, "/", index))
	// }
	 
	mux.HandleFunc(ENDPOINT_LLM_API, llm_api)
	mux.HandleFunc("/", index)
	

	API_PORT := os.Getenv(API_PORT_VARIABLE_NAME)
	if API_PORT == "" {
		logger.Error("you should set an environment variable: %s", API_PORT_VARIABLE_NAME)
	}

	logger.Info("start llm_api server on port: " + API_PORT)

	err = http.ListenAndServe(":"+API_PORT, mux)
	if err != nil {
		logger.Error("ListenAndServe: ", err)
	}
}
