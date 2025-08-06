package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

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

func (c *Client) WSReader(dcChan chan *Client) {

	for {

		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(c.ID, err)
			dcChan <- c
			return
		}
		fmt.Println(c.ID, string(msg))

	}

}

func (c *Client) WSWriter(dcChan chan *Client) {

	for {
		msg := <-c.ToClient
		err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			fmt.Println(c.ID, err)
			dcChan <- c
			return
		}
	}
}
