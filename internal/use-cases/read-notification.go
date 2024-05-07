package usecases

import (
	"github/eovinicius/notification/internal/repository"

	"github.com/google/uuid"
)

type ReadNotification struct {
	NotificationRepository repository.NotificationRepository
}

func NewReadNotification(notificationRepository repository.NotificationRepository) *ReadNotification {
	return &ReadNotification{
		NotificationRepository: notificationRepository,
	}
}

func (rn *ReadNotification) Execute(notificationID uuid.UUID) error {
	notification, err := rn.NotificationRepository.FindByID(notificationID)
	if err != nil {
		return err
	}

	notification.Read()

	err = rn.NotificationRepository.Save(notification)
	if err != nil {
		return err
	}

	return nil
}
