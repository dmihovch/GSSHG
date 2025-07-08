package game

import "github.com/gorilla/websocket"

type Player struct {
	Conn  *websocket.Conn
	chips int
	hand  []string
	name  string
}
