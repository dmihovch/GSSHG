package main

import (
	"log"
	"net/http"
	"sync"

	"gsshg/server/game"
	"gsshg/server/serversockets"
)

func main() {

	manager := &game.Manager{
		WaitGroup: &sync.WaitGroup{},
		ConnPool:  &game.ConnectionPool{},
	}
	manager.WaitGroup.Add(1)
	manager.Run()
	manager.WaitGroup.Wait()
	http.HandleFunc("/ws", serversockets.Handler)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
