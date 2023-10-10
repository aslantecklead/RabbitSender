package dependencies

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB

type MsgRecived struct {
	FromEmail    string `json:"From_Email"`
	ToEmail      string `json:"To_Email"`
	MsgTitle     string `json:"MsgTitle"`
	EmailBody    string `json:"EmailBody"`
	Timestamp    int    `json:"Timestamp"`
	ErrorMessage string `json:"Error_message"`
}

type sender struct {
	ID_Sender   int `gorm:"primary_key"`
	Sender_mail string
}

type emailmessage struct {
	ID_EmailMessage int `gorm:"primary_key"`
	email_title     string
	email_body      string
	reciver_mail    string // Исправлено на receiver_mail
}

type messageslogs struct {
	ID_MessagesLogs int   `gorm:"primary_key"`
	timestamp       int64 // Изменено на int64
	error_message   string
}

func InitDB() (*gorm.DB, error) {
	connStr := "host=localhost port=5433 dbname=msgSenderMicroservice user=postgres password=123 sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&sender{}, &emailmessage{}, &messageslogs{})

	db.SingularTable(true)
	return db, nil
}

func SaveToDb(d *gorm.DB, messageBytes []byte) {

	fmt.Printf("Получено сообщение: %s\n", messageBytes)

	var message MsgRecived
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		log.Printf("Ошибка разбора JSON: %v", err)
		return
	}

	senderRecord := sender{
		Sender_mail: message.FromEmail,
	}

	emailMessageRecord := emailmessage{
		email_title:  message.MsgTitle,
		email_body:   message.EmailBody,
		reciver_mail: message.ToEmail,
	}

	messageslogsRecord := messageslogs{
		timestamp:     int64(message.Timestamp),
		error_message: message.ErrorMessage,
	}

	if err := d.Create(&senderRecord).Error; err != nil {
		log.Printf("Ошибка при сохранении в таблице 'sender': %v", err)
	}

	if err := d.Create(&emailMessageRecord).Error; err != nil {
		log.Printf("Ошибка при сохранении в таблице 'emailmessage': %v", err)
	}

	if err := d.Create(&messageslogsRecord).Error; err != nil {
		log.Printf("Ошибка при сохранении в таблице 'messageslogs': %v", err)
	}
}
