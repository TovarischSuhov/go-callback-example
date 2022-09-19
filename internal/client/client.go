package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/TovarischSuhov/go-callback-example/internal/api"
)

var defaultClient http.Client

func SendMessage(messageName string, sleepTime int, id int) {
	msg := api.Message{Message: messageName, Sleep: sleepTime, ID: id}
	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := defaultClient.Post("http://localhost:8080/ping", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(r))
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Finish: '%s'", req)
}
