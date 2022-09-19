package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/TovarischSuhov/go-callback-example/internal/client"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go serve()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("Example%d", i)
		sleep := i % 3
		log.Printf("Send message '%s'\n", name)
		client.SendMessage(name, sleep, i)
	}
	wg.Wait()
}

func serve() {
	defer wg.Done()
	http.HandleFunc("/callback", client.CallbackHandler)
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Println(err)
	}
}
