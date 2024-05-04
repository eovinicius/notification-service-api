package usecases

import (
	"github/eovinicius/notification/internal/repository/interfaces"

	"github.com/google/uuid"
)

type CancelNotification struct {
	NotificationRepository interfaces.NotificationRepository
}

func NewCancelNotification(notificationRepository interfaces.NotificationRepository) *CancelNotification {
	return &CancelNotification{
		NotificationRepository: notificationRepository,
	}
}

func (cn *CancelNotification) Execute(notificationID uuid.UUID) error {

	notification, err := cn.NotificationRepository.FindByID(notificationID)
	if err != nil {
		return err
	}

	notification.Cancel()

	err = cn.NotificationRepository.Update(notification)
	if err != nil {
		return err
	}

	return nil
}
