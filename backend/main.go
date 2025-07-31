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
		ConnectionPool: make(map[string]*server.Client),
	}
	manager.WaitGroup.Add(1)
	go manager.AcceptConnections()
	manager.WaitGroup.Wait()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.Handler(w, r, manager)
	})
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
