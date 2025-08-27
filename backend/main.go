package main

import (
	"gsshg/backend/server"
	"log"
	"net/http"
)

func main() {

	manager := server.CreateManager()
	go manager.AcceptConnections()
	<-manager.Signals.ManagerReader

	go manager.MainGameLoop()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		manager.Handler(w, r)
	})
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
