package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func PublishMessage(message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"messages_queue", // имя очереди
		true,             // durable
		false,            // autoDelete
		false,            // exclusive
		false,            // noWait
		nil,              // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",           // exchange
		queue.Name,   // routing key (имя очереди)
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	log.Println("📨 Очередь объявлена:", queue.Name)
	log.Printf("✅ Сообщение отправлено в RabbitMQ: %s", message)
	return nil
}