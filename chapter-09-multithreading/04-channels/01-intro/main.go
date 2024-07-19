package main

import "fmt"

// Thread 1
func main() {

	channel := make(chan string) // empty channel

	// Thread 2
	go func() {
		channel <- "Hello World!" // full channel
	}()

	// Thread 1
	msg := <-channel // empty channel

	fmt.Println(msg)

}
