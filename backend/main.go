package main

import (
	"gsshg/backend/server"
	"log"
	"net/http"
	"sync"
)

func main() {

	manager := &server.Manager{
		Connections: &server.ConnectionPool{
			ConnMap:       make(map[int]*server.Client),
			IDarr:         []int{},
			SmallBlindID:  -1,
			CurrentTurnID: -1,
			Mutex:         &sync.Mutex{},
		},
		ClientChan:  make(chan (server.NewConnection)),
		ServerReady: make(chan (struct{})),
		StartGame:   make(chan (struct{})),
		GameState:   &server.GameState{},
	}
	go manager.AcceptConnections()
	<-manager.ServerReady

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		manager.Handler(w, r)
	})
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
