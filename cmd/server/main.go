package main

import (
	"log"
	"net/http"

	"github.com/TovarischSuhov/go-callback-example/internal/server"
)

func main() {
	http.HandleFunc("/ping", server.PingHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
