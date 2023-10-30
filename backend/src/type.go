package main

import "github.com/gorilla/websocket"

type Message struct {
	Message string `json:"message"`
	Created int64  `json:"created_at"`
}

type Request struct {
	Action  string `json:"action"`
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	Name    string `json:"name"`
}

type ErrorObject struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

type ChatResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Created int64  `json:"created_at"`
}

// WebSocket関連

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}
type subscription struct {
	conn *connection
	room string
}
type ByteBroadCast struct {
	Message []byte
	Type    int
	Conn    *websocket.Conn
}
type message struct {
	data []byte
	room string
}

type hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}
