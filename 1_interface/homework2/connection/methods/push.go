package methods

import (
	"Study_REST_API/1_interface/homework2/connection"
	"fmt"
)

type Push struct {
	Text string
}

func NewPush(text string) Push {
	return Push{
		Text: text,
	}
}

func (p Push) Name() string {
	return "Push-уведомление"
}

func (p Push) Send() {
	// Заглушка: имитация отправки push-уведомления
	fmt.Printf("Отправлено Push-уведомление: %s\n", p.Text)
}

func (p Push) GetInfo() connection.NotificationInfo {
	return connection.NewNotificationInfo(p.Name(), p.Text, "Отправлено")
}
