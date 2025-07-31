package game

import (
	"fmt"
	"sync"
)

type Manager struct {
	WaitGroup *sync.WaitGroup
	ConnPool  *ConnectionPool
}

func (m *Manager) Run() {
	defer m.WaitGroup.Done()
	fmt.Println("Manager Init")
}
