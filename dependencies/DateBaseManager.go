package dependencies

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

type MsgRecived struct {
	From      string
	ToEmail   string
	MsgTitle  string
	EmailBody string
	Timestamp int
}

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
	var err error
	db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
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

func InitDB() (*gorm.DB, error) {
	connStr := "host=localhost port=5433 dbname=msgSenderMicroservice user=postgres password=123 sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Sender{}, &EmailMessage{}, &MessagesLogs{})

	db.SingularTable(true)
	return db, nil
}

func SaveToDb(message MsgRecived) {

	fmt.Println(message)

	var emailMessages []EmailMessage
	if err := db.Find(&emailMessages).Error; err != nil {
		log.Printf("Не удалось выполнить SELECT-запрос: %v", err)
		return
	}
	/*
		sender := Sender{SenderMail: message.From}
		if err := db.Create(&sender).Error; err != nil {
			log.Printf("Не удалось создать запись в таблице 'sender': %v", err)
		}

		// Создать запись в таблице 'email_message'
		emailMessage := EmailMessage{
			EmailTitle:   message.MsgTitle,
			EmailBody:    message.EmailBody,
			ReceiverMail: message.ToEmail,
			IDSender:     sender.IDSender,
		}
		if err := db.Create(&emailMessage).Error; err != nil {
			log.Printf("Не удалось создать запись в таблице 'email_message': %v", err)
		}

		// Создать запись в таблице 'messages_logs'
		messagesLogs := MessagesLogs{
			Timestamp:      message.Timestamp,
			IDEmailMessage: emailMessage.IDEmailMessage,
		}
		if err := db.Create(&messagesLogs).Error; err != nil {
			log.Printf("Не удалось создать запись в таблице 'messages_logs': %v", err)
		}

		log.Println("Данные успешно сохранены в базе данных.")

	*/
}
