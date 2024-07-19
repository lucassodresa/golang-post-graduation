package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task: %s, Iteration: %d\n", name, i)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	// Thread 4
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Task: %s, Iteration: %d\n", "Anonymous", i)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
