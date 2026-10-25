package methods

import (
	"Study_REST_API/1_interface/homework2/connection"
	"fmt"
)

type SMS struct {
	Text string
}

func NewSMS(text string) SMS {
	return SMS{
		Text: text,
	}
}

func (s SMS) Name() string {
	return "SMS"
}

func (s SMS) Send() {
	// Заглушка: имитация отправки SMS
	fmt.Printf("Отправлено SMS: %s\n", s.Text)
}

func (s SMS) GetInfo() connection.NotificationInfo {
	return connection.NewNotificationInfo(s.Name(), s.Text, "Отправлено")
}
