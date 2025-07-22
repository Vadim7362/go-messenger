package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func StartConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("❌ Ошибка при открытии канала: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"messages_queue", // очередь та же, что и у продюсера
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Ошибка при объявлении очереди: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,  // авто-подтверждение
		false, // не эксклюзивный
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Ошибка при регистрации консьюмера: %v", err)
	}

	log.Println("📥 Консьюмер ожидает сообщений...")

	go func() {
		for d := range msgs {
			log.Printf("✅ Получено сообщение: %s", d.Body)
		}
	}()
}