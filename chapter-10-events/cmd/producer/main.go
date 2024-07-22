package main

import "github.com/lucassodresa/golang-post-graduation/chapter-10-events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, RabbitMQ!", "amq.direct")
}
