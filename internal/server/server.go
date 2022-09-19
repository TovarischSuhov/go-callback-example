package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/TovarischSuhov/go-callback-example/internal/api"
)

var defaultClient http.Client

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Not allowed method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var msg api.Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	go callbackUpdater(msg)
	w.WriteHeader(http.StatusOK)
}

func callbackUpdater(msg api.Message) {
	time.Sleep(time.Duration(msg.Sleep) * time.Second)
	resp := api.Response{Message: msg.Message, ID: msg.ID}
	buf, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = defaultClient.Post("http://localhost:8081/callback", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Println(err)
		return
	}
}
