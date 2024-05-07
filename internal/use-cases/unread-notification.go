package usecases

import (
	"github/eovinicius/notification/internal/repository"

	"github.com/google/uuid"
)

type UnreadNotification struct {
	NotificationRepository repository.NotificationRepository
}

func NewUnreadNotification(notificationRepository repository.NotificationRepository) *UnreadNotification {
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
