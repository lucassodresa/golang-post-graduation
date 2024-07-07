package main

import "github.com/lucassodresa/golang-post-graduation/chapter-08-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
