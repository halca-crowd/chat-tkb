package main

import (
	"log/slog"
	"os"
)

type HTTPLogger struct {
	infoLogHandler  slog.Handler // Infoレベルのログを出力するハンドラ
	errorLogHandler slog.Handler // Errorレベルのログを出力するハンドラ
	isInfo          bool         // Infoレベルのログを出力するか否か
}

// HTTPステータスコードに応じてログを出力する
func (hl *HTTPLogger) LoggingHTTPError(statusCode int, err error) {
	he := NewHTTPError(statusCode, err)

	switch he.Level {
	case Info:
		if hl.isInfo {
			infoLogger := slog.New(hl.infoLogHandler)
			infoLogger.Info(he.Error())
		}
	case Debug:
		debugLogger := slog.New(hl.infoLogHandler)
		debugLogger.Debug(he.Error())
	case Warning:
		warningLogger := slog.New(hl.infoLogHandler)
		warningLogger.Warn(he.Error())
	case Error:
		fatalLogger := slog.New(hl.errorLogHandler)
		fatalLogger.Error(he.Error())
		// When Error, exit
		os.Exit(1)
	}
}

// HTTPLoggerのコンストラクタ
func NewHTTPLogger(isInfo bool) *HTTPLogger {
	return &HTTPLogger{
		infoLogHandler:  slog.NewTextHandler(os.Stdout, nil),
		errorLogHandler: slog.NewTextHandler(os.Stderr, nil),
		isInfo:          isInfo}
}