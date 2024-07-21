package main

import "fmt"

func main() {
	events := []string{"test", "test2", "test3", "test4", "test5"}
	events = append(events[:0], events[1:]...)
	fmt.Println(events)

}
