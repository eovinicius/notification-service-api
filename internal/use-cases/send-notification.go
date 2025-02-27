package usecases

import (
	"github/eovinicius/notification/internal/entity"
	"github/eovinicius/notification/internal/repository"
	"github/eovinicius/notification/internal/services"

	"github.com/google/uuid"
)

type SendNotification struct {
	NotificationRepository repository.NotificationRepository
	EmailSender            services.EmailSender
}

func NewSendNotification(notificationRepository repository.NotificationRepository) *SendNotification {
	return &SendNotification{
		NotificationRepository: notificationRepository,
	}
}

func (sn *SendNotification) Execute(recipientID uuid.UUID, title, content, category string) error {

	_, err := sn.NotificationRepository.FindByID(recipientID)
	if err != nil {
		return err
	}

	notification, err := entity.NewNotification(recipientID, title, content, category)
	if err != nil {
		return err
	}

	err = sn.EmailSender.SendEmail(notification.RecipientID.String(), notification.Category, notification.Content)

	if err != nil {
		return err
	}

	err = sn.NotificationRepository.Save(notification)
	if err != nil {
		return err
	}

	return nil
}
