package main

import (
	"fmt"

	"github.com/lucassodresa/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
