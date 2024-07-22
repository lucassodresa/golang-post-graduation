package main

import (
	"fmt"

	"github.com/lucassodresa/utils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
