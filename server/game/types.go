package game

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	ToClient chan []byte
	Actions  chan Action
	State    *PlayerState
}

func CreateClient(conn *websocket.Conn) *Client {
	return &Client{
		ID:       uuid.New().String(),
		Conn:     conn,
		ToClient: make(chan []byte),
		Actions:  make(chan Action),
	}
}

type ConnectionPool struct {
	Players map[string]*PlayerState
}

type Action struct {
	Type    string
	Payload json.RawMessage
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
