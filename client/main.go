package main

import (
	"bufio"
	"fmt"
	"gsshg/client/game"
	"log"
	"os"
	"strconv"

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

	lnReader := bufio.NewReader(os.Stdin)

	_, msg, err := player.Conn.ReadMessage()
	if err != nil {
		return
	}

	fmt.Println(string(msg))

	username, err := lnReader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	player.Name = username
	player.Conn.WriteMessage(websocket.TextMessage, []byte(username))

	_, id, err := player.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	playerID, err := strconv.Atoi(string(id))
	if err != nil {
		log.Println(err)
		return
	}
	player.ID = playerID

	fmt.Println(player.ID, player.Name)

	for {

		clientMsg, err := lnReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if clientMsg == "q\n" {
			fmt.Println("disconnecting")

			break
		}

		if clientMsg == "w\n" {
			err = player.WriteJson(&game.JSONPayload{Type: "raise", Data: "200"})
			if err != nil {
				log.Fatalln(err)
			}
		}

		err = player.WriteTextMessage(clientMsg)
		if err != nil {
			log.Println(err)
			break
		}

	}

}
