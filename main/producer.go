package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
	"time"
)

type Message struct {
	From_Email    string
	To_Email      string
	MsgTitle      string
	EmailBody     string
	Timestamp     int64
	Error_message string // По умолчанию "0"
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

	fmt.Println("Введите сообщения. Для выхода введите 'exit'.")
	for {
		message := Message{}

		fmt.Print("Отправитель: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		message.From_Email = input

		fmt.Print("Получатель: ")
		scanner.Scan()
		message.To_Email = scanner.Text()

		fmt.Print("Тема письма: ")
		scanner.Scan()
		message.MsgTitle = scanner.Text()

		fmt.Print("Текст письма: ")
		scanner.Scan()
		message.EmailBody = scanner.Text()

		message.Timestamp = time.Now().Unix()

		// Обработка ошибки и установка значения поля Error_message
		var emptyFields []string
		if message.From_Email == "" {
			emptyFields = append(emptyFields, "Field 'Отправитель'")
		}
		if message.To_Email == "" {
			emptyFields = append(emptyFields, "Field 'Получатель'")
		}
		if message.MsgTitle == "" {
			emptyFields = append(emptyFields, "Field 'Тема письма'")
		}
		if message.EmailBody == "" {
			emptyFields = append(emptyFields, "Field 'Текст письма'")
		}

		if len(emptyFields) > 0 {
			message.Error_message = "Empty fields: " + strings.Join(emptyFields, ", ")
		} else if strings.Contains(message.EmailBody, "ошибка") {
			message.Error_message = "Обнаружена ошибка в тексте письма"
		} else {
			message.Error_message = "0"
		}

		jsonData, err := json.Marshal(message)
		if err != nil {
			log.Printf("Не удалось преобразовать в JSON: %v", err)
		} else {
			err = ch.Publish(
				"",
				queueName,
				false,
				false,
				amqp.Publishing{
					ContentType: "application/json",
					Body:        jsonData,
				},
			)
			if err != nil {
				log.Printf("Не удалось отправить сообщение: %v", err)
			} else {
				fmt.Println("Сообщение отправлено.")
			}
		}
	}
}
