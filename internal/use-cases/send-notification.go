package usecases

import (
	"github/eovinicius/notification/internal/entity"
	"github/eovinicius/notification/internal/repository/interfaces"

	"github.com/google/uuid"
)

type SendNotification struct {
	NotificationRepository interfaces.NotificationRepository
}

func NewSendNotification(notificationRepository interfaces.NotificationRepository) *SendNotification {
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
