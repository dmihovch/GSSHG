package main

import (
	"fmt"
	"gsshg/client/game"
	"log"

	"github.com/gorilla/websocket"
)

func main() {

	/*
		prog := tea.NewProgram(tui.CreateModel())
			if _, err := prog.Run(); err != nil {
				return
			}
	*/

	player := &game.Player{}
	player.ConnectToServer("ws://localhost:8080/ws")
	defer player.Conn.Close()

	err := player.Conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
	if err != nil {
		log.Fatalln(err)
	}
	_, msg, err := player.Conn.ReadMessage()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server says: " + string(msg))
}
