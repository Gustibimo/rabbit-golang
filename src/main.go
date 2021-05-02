package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go client()
	go server()

	var a string
	fmt.Scanln(&a) // to keep main go routine alive
}


func client()  {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err , "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Receive message with message: %s", msg.Body)
	}
}

func server()  {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte("Hello RabbitMQ"),
	}

	for
	{
		ch.Publish("", q.Name, false, false, msg)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect Rabbit")
	ch, err := conn.Channel()
	failOnError(err, "Failed to connect Channel")
	q, err := ch.QueueDeclare("hello", 
		false, 
		false,
		false,
		false,
		nil)

	failOnError(err, "failed to decalre queue")
	return conn, ch, &q

}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
