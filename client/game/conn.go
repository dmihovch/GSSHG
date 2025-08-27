package game

import (
	"fmt"
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

func (p *Player) WSReader() {
	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(msg))
	}
}

func (p *Player) WSWriter() {

	for {

		msg := <-p.ToServer
		err := p.WriteTextMessage(string(msg))
		if err != nil {
			log.Fatalln(err)
		}

	}
}
