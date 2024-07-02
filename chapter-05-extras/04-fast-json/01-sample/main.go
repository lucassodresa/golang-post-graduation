package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser

	jsonData := `{"foo": "bar", "num": 123, "bool": true, "arr": [1, 2, 3]}`

	parsedData, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("foo: %s\n", parsedData.GetStringBytes("foo"))
	fmt.Printf("num: %d\n", parsedData.GetInt("num"))
	fmt.Printf("bool: %t\n", parsedData.GetBool("bool"))

	array := parsedData.GetArray("arr")
	for i, value := range array {
		fmt.Printf("arr[%d]: %d\n", i, value.GetInt())
	}

}
