package server

import (
	"fmt"
	"sync"
)

type Manager struct {
	WaitGroup   *sync.WaitGroup
	Connections *ConnectionPool
	ClientChan  chan (NewConnection)
	Ready       chan (struct{})
}

func (m *Manager) AcceptConnections() {

	nextID := 0
	fmt.Println("accepting connections")
	close(m.Ready)
	for client := range m.ClientChan {
		m.Connections.Mutex.Lock()
		m.Connections.ConnMap[nextID] = CreateClient(client.conn, nextID, client.username)
		if len(m.Connections.ConnMap) == 1 {
			m.Connections.ConnMap[nextID].IsLeader = true
		}
		m.Connections.Mutex.Unlock()
		nextID++

		debugClient := m.Connections.ConnMap[nextID-1]
		fmt.Println(debugClient.ID, debugClient.ScreenName)

	}

}
