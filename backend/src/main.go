package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"*",
		},
		MaxAge: 24 * time.Hour,
	}))
	// mux := http.NewServeMux()
	//キャッシュの作成と導通テスト（他のやり方があれば直す）
	cacheManager := cached.NewMemcached("memcached:11211")
	//1分だけ有効なテスト値を挿入する
	cacheManager.SaveFor(time.Minute, "test", []byte("test"))
	testValue, err := cacheManager.Get("test")
	if len(testValue) == 0 || err != nil {
		log.Fatalln("failed to connect memcached")
	}

	// router.GET("/llm_api", func(c *gin.Context) {
	// 	llm_api(c.Writer, c.Request)
	// })

	// ChatTBKの受信用

	router.GET("/ws", func(c *gin.Context) {
		// roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, "last-hack")
	})

	// プリセットの取得用API

	router.GET("/preset", func(c *gin.Context) {
		preset(c.Writer, c.Request)
	})
	// ヘルスチェック用
	router.GET("/", func(c *gin.Context) {
		index(c.Writer, c.Request)
	})
	// 管理用（Redisの貯まっているデータを強制削除する）
	router.POST("/reset", func(c *gin.Context) {
		reset(c.Writer, c.Request)
	})

	// 管理用（GETでLLMにプロンプトを投げる）
	router.GET("/llm_api", func(c *gin.Context) {
		llm_api(c.Writer, c.Request)
	})

	API_PORT := os.Getenv(API_PORT_VARIABLE_NAME)
	if API_PORT == "" {
		logger.Error("you should set an environment variable: %s", API_PORT_VARIABLE_NAME)
	}

	logger.Info("start llm_api server on port: " + API_PORT)
	router.Run(":8080")
	// err = http.ListenAndServe(":"+API_PORT, mux)
	if err != nil {
		logger.Error("ListenAndServe: ", err)
	}
}
