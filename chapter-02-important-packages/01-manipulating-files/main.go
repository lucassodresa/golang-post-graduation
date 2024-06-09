package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	size, err := file.Write([]byte("The house blue is the most beautiful in the househood."))
	// size, err := file.WriteString("Hey man")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File created with success! The size is: %d bytes\n", size)

	file.Close()

	// read file
	fileRead, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileRead))

	// read file line by line
	fileByLine, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileByLine)
	buffer := make([]byte, 10)

	for {

		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// remove file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
