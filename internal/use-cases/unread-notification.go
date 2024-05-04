package usecases

import (
	"github/eovinicius/notification/internal/repository/interfaces"

	"github.com/google/uuid"
)

type UnreadNotification struct {
	NotificationRepository interfaces.NotificationRepository
}

func NewUnreadNotification(notificationRepository interfaces.NotificationRepository) *UnreadNotification {
	return &UnreadNotification{
		NotificationRepository: notificationRepository,
	}
}

func (un *UnreadNotification) Execute(recipientID uuid.UUID) error {
	notification, err := un.NotificationRepository.FindByID(recipientID)
	if err != nil {
		return err
	}

	notification.Unread()

	err = un.NotificationRepository.Save(notification)
	if err != nil {
		return err
	}

	return nil
}
