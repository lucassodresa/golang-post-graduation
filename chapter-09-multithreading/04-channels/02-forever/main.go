package main

// Thread 1
func main() {
	forever := make(chan bool) // empty channel

	// Thread 2
	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // full channel
	}()

	<-forever // empty channel

}
