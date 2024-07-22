package main

import (
	"fmt"

	"github.com/lucassodresa/golang-post-graduation/chapter-10-events/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
