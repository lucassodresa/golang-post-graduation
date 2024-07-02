package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var p fastjson.Parser
	jsonData := `{"user": { "name": "John Doe", "age": 30}}`

	parsedData, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJSON := parsedData.Get("user").String()

	var user User

	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}

	fmt.Println(user.Name, user.Age)

	// fmt.Printf("Name: %s\n", user.Get("name"))
	// fmt.Printf("Age: %s\n", user.Get("age"))

}
