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
	ToClient   chan []byte //or this
	Actions    chan Action //don't think I need this
	State      *PlayerState
	IsLeader   bool
	IsTurn     bool
}

type Manager struct {
	Connections *ConnectionPool
	ClientChan  chan (NewConnection)
	ServerReady chan (struct{})
	StartGame   chan (struct{})
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
	ConnMap       map[int]*Client
	IDarr         []int
	SmallBlindID  int
	CurrentTurnID int
	Mutex         *sync.Mutex
}

type JSONPayload struct {
	Type string `json:"type"` //type determines which handler is used (eg raise, fold, check, flip)
	Data string `json:"data"` //determines quantity, whatever else.
}
