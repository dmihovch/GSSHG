package game

import "github.com/gorilla/websocket"

type Player struct {
	Conn     *websocket.Conn
	ToServer chan ([]byte)
	Chips    int
	Hand     []string
	Name     string
	ID       int
}

//going to have a list of keybind, if a key is pressed, send message to the server with keypress and current game state (tag each gamestate with an id? idk yet),
// process the request

type JSONPayload struct {
	Type string `json:"type"` //type determines which handler is used (eg raise, fold, check, flip)
	Data string `json:"data"` //determines quantity, whatever else.
}
