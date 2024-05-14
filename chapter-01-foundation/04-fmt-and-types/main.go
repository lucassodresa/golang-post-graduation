package main

import "fmt"

type ID int

var (
	userID ID = 1
)

func main() {
	fmt.Printf("The type of decimal is %T", userID)
}
