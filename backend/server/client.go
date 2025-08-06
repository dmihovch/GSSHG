package server

import (
	"fmt"
	"time"

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

		msgType, msg, err := c.Conn.ReadMessage()
		if err != nil {
			dcChan <- c
		}
		fmt.Println(msgType)
		fmt.Println(string(msg))

	}

}

func (c *Client) WSWriter(dcChan chan *Client) {

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case msg, ok := <-c.ToClient:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			err := c.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				dcChan <- c
				return
			}

		case <-ticker.C:
			fmt.Println("hello") //disconnect
		}
	}
}
