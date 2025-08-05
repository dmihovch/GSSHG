package server

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func (m *Manager) AcceptConnections() {

	nextID := 0
	fmt.Println("accepting connections")
	close(m.ServerReady)
	for client := range m.ClientChan {
		m.Connections.Mutex.Lock()
		m.Connections.ConnMap[nextID] = CreateClient(client.conn, nextID, client.username)
		if len(m.Connections.ConnMap) == 1 {
			m.Connections.ConnMap[nextID].IsHost = true
			m.Connections.ConnMap[nextID].IsTurn = true
			m.Connections.SmallBlindID = nextID
			m.Connections.CurrentTurnID = nextID
		}
		m.Connections.IDarr = append(m.Connections.IDarr, nextID)
		m.Connections.Mutex.Unlock()
		nextID++

		debugClient := m.Connections.ConnMap[nextID-1]
		fmt.Println(debugClient.ID, debugClient.ScreenName)

	}

}

func (m *Manager) MainGameLoop() {

	m.ResetGameState()

}

func CreateClient(cn *websocket.Conn, id int, un string) *Client {
	return &Client{
		ScreenName: un,
		ID:         id,
		Conn:       cn,
		ToClient:   make(chan ([]byte)),
		Actions:    make(chan (Action)),
		State:      &PlayerState{},
	}
}
