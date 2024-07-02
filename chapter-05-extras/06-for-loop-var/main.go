package main

import "fmt"

func main() {

	for i := range 10 {
		fmt.Println(i)

	}

	done := make(chan bool)
	values := []string{"a", "b", "c"}

	for _, value := range values {
		go func() {
			fmt.Println(value)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
