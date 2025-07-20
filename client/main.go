package main

import (
	"bufio"
	"fmt"
	"gsshg/client/game"
	"log"
	"os"
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

	for {

		msg, err := lnReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		if msg == "q\n" {
			fmt.Println("disconnecting")
			break
		}

		if msg == "w\n" {
			err = player.WriteJson(&game.JSONPayload{Type: "raise", Data: "200"})
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

}
