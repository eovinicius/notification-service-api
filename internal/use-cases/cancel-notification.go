package usecases

import (
	"github/eovinicius/notification/internal/repository"

	"github.com/google/uuid"
)

type CancelNotification struct {
	NotificationRepository repository.NotificationRepository
}

func NewCancelNotification(notificationRepository repository.NotificationRepository) *CancelNotification {
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
