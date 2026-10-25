package connection

// NotificationInfo структура для хранения информации об уведомлении
type NotificationInfo struct {
	Type   string // Тип уведомления (Email, SMS, Push)
	Text   string // Текст уведомления
	Status string // Статус отправки (например, "Отправлено", "Ошибка")
}

// NewNotificationInfo создаёт новую структуру информации
func NewNotificationInfo(notificationType, text, status string) NotificationInfo {
	return NotificationInfo{
		Type:   notificationType,
		Text:   text,
		Status: status,
	}
}
