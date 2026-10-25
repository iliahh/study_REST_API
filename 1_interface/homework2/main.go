package main

import (
	"Study_REST_API/1_interface/homework2/connection"
	"Study_REST_API/1_interface/homework2/connection/methods"
)

func main() {
	// Создаём уведомления разных типов
	email := methods.NewEmail("Привет! Это тестовое email.")
	sms := methods.NewSMS("Привет! Это тестовое SMS.")
	push := methods.NewPush("Привет! Это тестовое push-уведомление.")

	// Срез уведомлений
	notifications := []connection.NotificationSender{
		email,
		sms,
		push,
	}

	// Отправляем все уведомления
	connection.SendNotifications(notifications)

	// Выводим информацию по уведомлениям
	connection.PrintNotificationsInfo(notifications)
}
