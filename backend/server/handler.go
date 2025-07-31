package server

import (
	"fmt"
	"gsshg/global"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request, m *Manager) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection added:", conn.LocalAddr().String())

	for {
		/*
			 _, msg, err := conn.ReadMessage()
				if err != nil {
					log.Println(err)
					break
				}

				fmt.Println("msg from client: " + "{" + string(msg) + "}")

				err = conn.WriteMessage(websocket.TextMessage, []byte("Hey idiot: "+string(msg)))
				if err != nil {
					log.Println(err)
					break
				}
		*/

		var payload global.JSONPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("Client disconnected:", conn.LocalAddr().Network())
				break
			}
			fmt.Println(err)
			break
		}

		fmt.Println("Payload:")
		fmt.Println(payload.Type)
		fmt.Println(payload.Data)

	}

}
