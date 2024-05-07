package usecases

import (
	"github/eovinicius/notification/internal/entity"
	"github/eovinicius/notification/internal/repository"
)

type GetRecipientNotifications struct {
	NotificationRepository repository.NotificationRepository
}

func NewGetRecipientNotifications(notificationRepository repository.NotificationRepository) *GetRecipientNotifications {
	return &GetRecipientNotifications{
		NotificationRepository: notificationRepository,
	}
}

func (grn *GetRecipientNotifications) Execute(recipientID string) ([]*entity.Notification, error) {
	notifications, err := grn.NotificationRepository.FindManyByRecipientID(recipientID)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
