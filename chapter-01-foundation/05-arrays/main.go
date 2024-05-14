package main

import "fmt"

type ID int

func main() {
	var array [3]int

	array[0] = 10
	array[1] = 20
	array[2] = 30

	lastPosition := len(array) - 1
	size := len(array)
	lastItem := array[lastPosition]

	fmt.Println(lastPosition, size, lastItem)

	for index, value := range array {
		fmt.Printf("The value of index %d is %d\n", index, value)
	}
}
