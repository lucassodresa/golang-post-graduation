package main

import "fmt"

func receive(name string, ch chan<- string) {
	ch <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	ch := make(chan string)
	go receive("Lucas", ch)
	read(ch)
}
