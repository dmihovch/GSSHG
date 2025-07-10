package game

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      string
	Conn    *websocket.Conn
	Send    chan []byte
	Actions chan Action
}

func CreateClient(conn *websocket.Conn) *Client {
	return &Client{
		ID:      uuid.New().String(),
		Conn:    conn,
		Send:    make(chan []byte),
		Actions: make(chan Action),
	}
}

type GameState struct {
	Players map[string]*PlayerState
}

type Action struct {
	PlayerID string
	Type     string
	Payload  json.RawMessage
}

type PlayerState struct {
	Hand   [7]*Card
	Chips  int
	Folded bool
}

type Card struct {
	Suit  int
	Value int
}
