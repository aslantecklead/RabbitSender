package dependencies

import (
	"encoding/json"
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
	EmailTitle   string
	EmailBody    string
	ReceiverMail string
	ID_Sender    int
}

type messageslogs struct {
	ID_MessagesLogs int `gorm:"primary_key"`
	timestamp       int64
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
	//fmt.Printf("Получено сообщение: %s\n", messageBytes)

	var message MsgRecived
	if err := json.Unmarshal(messageBytes, &message); err != nil {
		log.Printf("Ошибка разбора JSON: %v", err)
		return
	}

	// Начать транзакцию
	tx := d.Begin()
	if tx.Error != nil {
		log.Printf("Ошибка при начале транзакции: %v", tx.Error)
		return
	}

	senderRecord := sender{
		Sender_mail: message.FromEmail,
	}

	if err := tx.Create(&senderRecord).Error; err != nil {
		// Откатить транзакцию в случае ошибки
		tx.Rollback()
		log.Printf("Ошибка при сохранении в таблице 'sender': %v", err)
		return
	}

	senderID := senderRecord.ID_Sender

	emailMessageRecord := emailmessage{
		EmailTitle:   message.MsgTitle,
		EmailBody:    message.EmailBody,
		ReceiverMail: message.ToEmail,
		ID_Sender:    senderID, // Используйте полученный ID отправителя
	}

	if err := tx.Create(&emailMessageRecord).Error; err != nil {
		// Откатить транзакцию в случае ошибки
		tx.Rollback()
		log.Printf("Ошибка при сохранении в таблице 'emailmessage': %v", err)
		return
	}

	messageslogsRecord := messageslogs{
		timestamp:     int64(message.Timestamp),
		error_message: message.ErrorMessage,
	}

	if err := tx.Create(&messageslogsRecord).Error; err != nil {
		tx.Rollback()
		log.Printf("Ошибка при сохранении в таблице 'messageslogs': %v", err)
		return
	}

	// Завершить транзакцию
	tx.Commit()
	SendEmail(messageBytes)
}
