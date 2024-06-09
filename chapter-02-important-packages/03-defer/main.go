package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First Line")
	defer fmt.Println("Second Line")
	fmt.Println("Third Line")
}
