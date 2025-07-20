package game

import (
	"log"

	"github.com/gorilla/websocket"
)

func (p *Player) ConnectToServer(url string) error {

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	p.Conn = conn
	return nil
}

func (p *Player) WriteTextMessage(msg string) error {
	return p.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (p *Player) WriteJson(payload *JSONPayload) error {
	return p.Conn.WriteJSON(payload)
}
