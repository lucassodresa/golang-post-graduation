package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := sum(51, 2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func sum(a int, b int) (int, error) {
	if a+b > 50 {
		return a + b, errors.New("Sum is greater than 50")
	}

	return a + b, nil
}
