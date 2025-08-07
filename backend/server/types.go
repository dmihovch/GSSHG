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
	IsHost     bool
	IsTurn     bool
}

type Manager struct {
	Connections      *ConnectionPool
	ClientChan       chan (NewConnection)
	GameState        *GameState
	DisconnectClient chan (*Client)
	Signals          *SignalChannels
}

type GameState struct {
	Started        bool
	TurnID         int
	TurnIndex      int
	CommunityCards [5]*Card
	Deck           [52]*Card
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
	Suit        int
	Value       int
	IsCommunity bool
	IsRevealed  bool
}

type NewConnection struct {
	conn     *websocket.Conn
	username string
}

type ConnectionPool struct {
	ConnMap       map[int]*Client
	IDarr         []int //determines the order of the game
	SmallBlindID  int
	CurrentTurnID int
	Mutex         *sync.Mutex
}

type JSONPayload struct {
	Type string `json:"type"` //type determines which handler is used (eg raise, fold, check, flip)
	Data string `json:"data"` //determines quantity, whatever else.
}

type SignalChannels struct {
	ManagerReader chan (struct{})
	StartGame     chan (struct{})
}
