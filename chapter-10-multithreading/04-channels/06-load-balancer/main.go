package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for value := range data {
		fmt.Printf("Worker %d got %d\n", workerId, value)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	nWorkers := 10

	for i := 0; i < nWorkers; i++ {
		go worker(nWorkers, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}
