package connection

import "fmt"

// NotificationSender интерфейс для отправки уведомлений
type NotificationSender interface {
	Name() string
	Send()
	GetInfo() NotificationInfo
}

// SendNotifications принимает срез уведомлений и отправляет их
func SendNotifications(notifications []NotificationSender) {
	for _, notification := range notifications {
		notification.Send()
	}
}

// PrintNotificationsInfo выводит информацию по всем уведомлениям
func PrintNotificationsInfo(notifications []NotificationSender) {
	fmt.Println("=== Информация по уведомлениям ===")
	for i, notification := range notifications {
		info := notification.GetInfo()
		fmt.Printf("%d. Тип: %s, Текст: %s, Статус: %s\n",
			i+1, info.Type, info.Text, info.Status)
	}
}
