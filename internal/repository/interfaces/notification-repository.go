package interfaces

import (
	"github/eovinicius/notification/internal/entity"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	Save(notification *entity.Notification) error
	FindByID(id uuid.UUID) (*entity.Notification, error)
	Update(notification *entity.Notification) error
	FindManyByRecipientID(recipientID string) ([]*entity.Notification, error)
}
