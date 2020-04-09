package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

/*
Helper function
*/
func failOnError(err error, msg string) {

	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToRabbitMq() *amqp.Connection {

	// create a connection -
	// connect to a rabbitmq server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func main() {

	conn := connectToRabbitMq()
	// this means that the connection will be closed later,
	// after all things
	defer conn.Close()

	fmt.Println("Connected")

	// Next we need to create a channel
	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// then we declare a queue in order to send a message to it
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	// then we define the body of the message:
	body := "Hello World!"

	// and finally publish it
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")

	fmt.Println("Message published")
}
