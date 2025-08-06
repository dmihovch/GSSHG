package server

import (
	"fmt"
	"strconv"

	"github.com/gorilla/websocket"
)

func (m *Manager) AcceptConnections() {

	nextID := 0
	fmt.Println("accepting connections")
	close(m.ServerReady)
	for client := range m.ClientChan {
		m.Connections.Mutex.Lock()
		newClient := CreateClient(client.conn, nextID, client.username)
		m.Connections.ConnMap[nextID] = newClient
		if len(m.Connections.ConnMap) == 1 {
			m.Connections.ConnMap[nextID].IsHost = true
			m.Connections.ConnMap[nextID].IsTurn = true
			m.Connections.SmallBlindID = nextID
			m.Connections.CurrentTurnID = nextID
		}
		m.Connections.IDarr = append(m.Connections.IDarr, nextID)
		m.Connections.Mutex.Unlock()
		nextID++

		err := newClient.Conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(nextID-1)))
		if err != nil {
			m.DisconnectClient <- newClient
			fmt.Println(newClient.ID, err)
		}

		go newClient.WSReader(m.DisconnectClient)
		go newClient.WSWriter(m.DisconnectClient)

	}

}

func (m *Manager) MainGameLoop() {

}
