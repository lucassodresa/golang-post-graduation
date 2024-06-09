package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"s"`
}

func main() {
	account := Account{Number: 1, Balance: 100}

	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	rawJson := []byte(`{"n":2,"s":200}`)
	var accountX Account
	err = json.Unmarshal(rawJson, &accountX)
	if err != nil {
		println(err)
	}
	println(accountX.Balance)
}
