package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	client := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}
