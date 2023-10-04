package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

type Sender struct {
	IDSender   int
	SenderMail string
}

type EmailMessage struct {
	IDEmailMessage int
	EmailTitle     string
	EmailBody      string
	ReceiverMail   string
	IDSender       int
}

type MessagesLogs struct {
	IDMessagesLogs int
	Timestamp      int
	ErrorMessage   string
	IDEmailMessage int
}

func main() {
	connStr := "host=localhost port=5433 dbname=msgSenderMicroservice user=postgres password=123 sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	db.SingularTable(true)

	var senders []Sender
	db.Find(&senders)
	fmt.Println("Senders:")
	fmt.Println(senders)

	var emailMessages []EmailMessage
	db.Find(&emailMessages)
	fmt.Println("Email Messages:")
	fmt.Println(emailMessages)

	var messagesLogs []MessagesLogs
	db.Find(&messagesLogs)
	fmt.Println("Messages Logs:")
	fmt.Println(messagesLogs)
}

func saveToDb(message MessageRecived) {
	emailMessage := EmailMessage{
		EmailTitle:   message.MsgTitle,
		EmailBody:    message.EmailBody,
		ReceiverMail: message.ToEmail,
		IDSender:     0,
	}

	if err := db.Create(&emailMessage).Error; err != nil {
		log.Printf("Не удалось сохранить данные в базе данных: %v", err)
		return
	}

	log.Println("Данные успешно сохранены в базе данных.")
}
