package server

import (
	"fmt"
	"sync"
)

type Manager struct {
	WaitGroup      *sync.WaitGroup
	Connections    *ConnectionPool
	ConnectionPool map[int]*Client
	ClientChan     chan (NewConnection)
}

func (m *Manager) AcceptConnections() {

	nextID := 0
	fmt.Println("accepting connections")
	for client := range m.ClientChan {
		fmt.Println("adding client")
		m.Connections.Mutex.Lock()
		m.ConnectionPool[nextID] = CreateClient(client.conn, nextID, client.username)
		m.Connections.Mutex.Unlock()
		nextID++

		debugClient := m.ConnectionPool[nextID-1]
		fmt.Println(debugClient.ID, debugClient.ScreenName)

	}

}
