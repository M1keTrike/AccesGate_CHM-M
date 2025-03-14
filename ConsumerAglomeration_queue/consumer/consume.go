package consumer

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/sender"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(conn *amqp091.Connection, failOnError func(error, string)) {
	cq := os.Getenv("CONSUME_QUEUE")
	apiURL := os.Getenv("APIHEX_RETURNER_URL")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		cq,
		"ConsumerAPIReturner",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	sender := sender.NewSender(apiURL)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			sender.Send(d.Body)
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
