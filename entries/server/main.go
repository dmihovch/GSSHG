package main

import (
	"log"
	"net/http"

	"gsshg/server/serversockets"
)

func main() {

	http.HandleFunc("/ws", serversockets.Handler)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
