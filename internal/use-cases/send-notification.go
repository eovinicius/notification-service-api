package usecases

import (
	"github/eovinicius/notification/internal/entity"
	"github/eovinicius/notification/internal/repository"

	"github.com/google/uuid"
)

type SendNotification struct {
	NotificationRepository repository.NotificationRepository
}

func NewSendNotification(notificationRepository repository.NotificationRepository) *SendNotification {
	return &SendNotification{
		NotificationRepository: notificationRepository,
	}
}

func (sn *SendNotification) Execute(recipientID uuid.UUID, content, category string) error {

	_, err := sn.NotificationRepository.FindByID(recipientID)
	if err != nil {
		return err
	}

	notification, err := entity.NewNotification(recipientID, content, category)
	if err != nil {
		return err
	}

	err = sn.NotificationRepository.Save(notification)
	if err != nil {
		return err
	}

	return nil
}
