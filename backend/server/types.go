package server

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ScreenName string
	ID         int
	Conn       *websocket.Conn
	ToClient   chan []byte
	Actions    chan Action
	State      *PlayerState
	IsLeader   bool
	IsTurn     bool
}

func CreateClient(conn *websocket.Conn, id int, username string) *Client {
	return &Client{
		ScreenName: username,
		ID:         id,
		Conn:       conn,
		ToClient:   make(chan []byte),
		Actions:    make(chan Action),
		State: &PlayerState{
			Hand:   [7]*Card{},
			Chips:  1000,
			Folded: false,
		},
	}
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

type NewConnection struct {
	conn     *websocket.Conn
	username string
}

type ConnectionPool struct {
	ConnMap map[int]*Client
	Mutex   *sync.Mutex
}
