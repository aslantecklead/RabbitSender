package main

import (
	"RabbitsSender/dependencies"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type SenderObj struct {
	FromEmail string `json:"From_Email"`
}

type MessageRecived struct {
	ToEmail   string `json:"To_Email"`
	MsgTitle  string `json:"MsgTitle"`
	EmailBody string `json:"EmailBody"`
	Timestamp int64  `json:"Timestamp"`
}

type MessagesLogsObj struct {
	Timestamp    int    `json:"Timestamp"`
	ErrorMessage string `json:"Error_message"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Не удалось подключиться к RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Не удалось открыть канал: %v", err)
		return
	}
	defer ch.Close()

	queueName := "messages"

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Не удалось объявить очередь: %v", err)
		return
	}

	fmt.Printf("Ожидание сообщений в очереди %s. Для выхода, нажмите CTRL+C\n", queueName)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Не удалось зарегистрировать потребителя: %v", err)
		return
	}

	for msg := range msgs {
		fmt.Printf("Получено сообщение: %s\n", msg.Body)

		var messageRecived dependencies.MsgRecived
		if err := json.Unmarshal(msg.Body, &messageRecived); err != nil {
			log.Printf("Не удалось декодировать JSON: %v", err)
			continue
		}
		dependencies.SaveToDb(messageRecived)
	}
}
