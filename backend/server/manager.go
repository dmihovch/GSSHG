package server

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

func CreateManager() *Manager {
	return &Manager{
		Connections: &ConnectionPool{
			ConnMap:       make(map[int]*Client),
			IDarr:         []int{},
			SmallBlindID:  -1,
			CurrentTurnID: -1,
			Mutex:         &sync.Mutex{},
		},
		ClientChan: make(chan (NewConnection)),
		Signals: &SignalChannels{
			ManagerReader: make(chan (struct{})),
			StartGame:     make(chan (struct{})),
		},
		GameState:        &GameState{},
		DisconnectClient: make(chan (*Client)),
	}
}

func (m *Manager) AcceptConnections() {

	nextID := 0
	fmt.Println("accepting connections")
	close(m.Signals.ManagerReader)
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

		err := newClient.Conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(nextID)))
		if err != nil {
			m.DisconnectClient <- newClient
			fmt.Println(newClient.ID, err)
		}

		nextID++
		go newClient.WSReader(m.DisconnectClient)
		go newClient.WSWriter(m.DisconnectClient)

	}

}

func (m *Manager) MainGameLoop() {

}
