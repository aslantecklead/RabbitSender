package dependencies

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendEmail(msg []byte) {
	fmt.Printf("Сообщение: %s\n", msg)

	var message MsgRecived
	if err := json.Unmarshal(msg, &message); err != nil {
		fmt.Printf("Ошибка разбора JSON: %v\n", err)
		return
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", message.FromEmail)
	mailer.SetHeader("To", message.ToEmail)
	mailer.SetHeader("Subject", message.MsgTitle+": <"+message.FromEmail+">")
	mailer.SetBody("text/plain", message.EmailBody)

	//Тута (1)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "asalbekovaslan@gmail.com", "wvak hpld srkk uvwv")

	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Printf("Ошибка при отправке письма: %v\n", err)
	}
}
