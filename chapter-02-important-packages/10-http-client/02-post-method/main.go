package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {

	client := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"John Doe"}`))
	resp, err := client.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

}
