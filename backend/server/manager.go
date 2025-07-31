package server

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	WaitGroup      *sync.WaitGroup
	ConnectionPool map[string]*Client
	ConnChan       chan (*websocket.Conn)
}

func (m *Manager) AcceptConnections() {

	for conn := range m.ConnChan {

	}

	fmt.Println("Manager Init")
	m.WaitGroup.Done()
}
