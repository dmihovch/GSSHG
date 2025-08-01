package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *Manager) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//defer conn.Close()
	fmt.Println("Connection added:", conn.LocalAddr().String())

	err = conn.WriteMessage(websocket.TextMessage, []byte("Username: "))
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("recieved username:", string(msg))

	client := NewConnection{
		conn:     conn,
		username: string(msg),
	}

	fmt.Println(client)

	m.ClientChan <- client
	fmt.Println("sent client to channel")
}

/*


	for {

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
*/
