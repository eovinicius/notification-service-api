package usecases

import (
	"github/eovinicius/notification/internal/entity"
	"github/eovinicius/notification/internal/repository/interfaces"
)

type GetRecipientNotifications struct {
	NotificationRepository interfaces.NotificationRepository
}

func NewGetRecipientNotifications(notificationRepository interfaces.NotificationRepository) *GetRecipientNotifications {
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
