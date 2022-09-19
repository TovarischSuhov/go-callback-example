package api

type Message struct {
	Sleep   int    `json:"sleep"`
	Message string `json:"message"`
	ID      int    `json:"id"`
}

type Response struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}
