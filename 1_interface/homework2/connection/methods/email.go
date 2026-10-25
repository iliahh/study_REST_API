package methods

import (
	"Study_REST_API/1_interface/homework2/connection"
	"fmt"
)

type Email struct {
	Text string
}

func NewEmail(text string) Email {
	return Email{
		Text: text,
	}
}

func (c Email) Name() string {
	return "Email"
}

func (c Email) Send() {
	// Заглушка: имитация отправки email
	fmt.Printf("Отправлено Email: %s\n", c.Text)
}

func (c Email) GetInfo() connection.NotificationInfo {
	return connection.NewNotificationInfo(c.Name(), c.Text, "Отправлено")
}
