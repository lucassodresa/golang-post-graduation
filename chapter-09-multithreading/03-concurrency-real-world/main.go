package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		message := fmt.Sprintf("You are the visitor number: %d", number)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(message))
	})

	http.ListenAndServe(":3000", nil)
}
