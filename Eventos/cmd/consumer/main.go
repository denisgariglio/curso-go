package main

import (
	"fmt"

	"github.com/denisgariglio/fcutils/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // --> não colocar a mensagem novamente na fila
	}

}
