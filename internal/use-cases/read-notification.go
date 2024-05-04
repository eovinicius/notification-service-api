package usecases

import (
	"github/eovinicius/notification/internal/repository/interfaces"

	"github.com/google/uuid"
)

type ReadNotification struct {
	NotificationRepository interfaces.NotificationRepository
}

func NewReadNotification(notificationRepository interfaces.NotificationRepository) *ReadNotification {
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
