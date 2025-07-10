package main

import (
	"bufio"
	"fmt"
	"gsshg/client/game"
	"log"
	"os"

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
	err := player.ConnectToServer("ws://localhost:8080/ws")
	if err != nil {
		log.Fatalln(err)
	}
	defer player.Conn.Close()

	err = player.Conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
	if err != nil {
		log.Fatalln(err)
	}
	_, msg, err := player.Conn.ReadMessage()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server says: " + string(msg))

	lnReader := bufio.NewReader(os.Stdin)

	for {

		msg, err := lnReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if msg == "q\n" {
			fmt.Println("disconnecting")
			break
		}

		err = player.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Fatalln(err)
		}
	}

}
