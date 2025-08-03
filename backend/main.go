package main

import (
	"gsshg/backend/server"
	"sync"
)

func main() {

	manager := &server.Manager{
		WaitGroup: &sync.WaitGroup{},
		Connections: &server.ConnectionPool{
			ConnMap: make(map[int]*server.Client),
			Mutex:   &sync.Mutex{},
		},
		ClientChan: make(chan (server.NewConnection)),
		Ready:      make(chan (struct{})),
	}
	go manager.AcceptConnections()
	<-manager.Ready
	go manager.StartServer()
}
