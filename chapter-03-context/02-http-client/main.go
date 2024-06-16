package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		responseMessage := "Request proceed with success"
		log.Println(responseMessage)
		w.Write([]byte(responseMessage))
	case <-ctx.Done():
		log.Println("Request cancelled by the client")
	}
}
