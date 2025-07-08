package serversockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte("Hey idiot: "+string(msg)))
		if err != nil {
			log.Println(err)
			break
		}

	}

}
