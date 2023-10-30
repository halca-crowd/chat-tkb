package main

import (
	"fmt"
	"net/http"
)

// errorのinterfaceを満たす専用エラー型
type HTTPError struct {
	StatusCode    int
	Level         ErrorLevel
	internalError error
}

func (he HTTPError) Error() string {
	return fmt.Sprintf(
		"Level: %s, Status: %d %s, Message: %s",
		he.Level.String(),
		he.StatusCode,
		http.StatusText(he.StatusCode),
		he.internalError.Error(),
	)
}

func (he HTTPError) Unwrap() error {
	return he.internalError
}

// HTTPErrorのコンストラクタ
func NewHTTPError(statusCode int, err error) *HTTPError {
	e := new(HTTPError)
	e.internalError = err
	e.StatusCode = statusCode
	switch int(statusCode / 100) {
	case 2:
		e.Level = Info
	case 4:
		e.Level = Info
	case 5:
		e.Level = Warning
	default:
		e.Level = Info
	}
	return e
}