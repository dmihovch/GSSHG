package main

import (
	"gsshg/backend/server"
	"log"
	"net/http"
	"sync"
)

func main() {

	manager := &server.Manager{
		WaitGroup:      &sync.WaitGroup{},
		ConnectionPool: make(map[int]*server.Client),
		ClientChan:     make(chan (server.NewConnection)),
	}

	ready := make(chan (struct{}))
	go func() {
		close(ready)
		manager.AcceptConnections()
	}()
	<-ready
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		manager.Handler(w, r)
	})
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
