package repository

import (
	"database/sql"
	"github/eovinicius/notification/internal/entity"
)

type NotificationRepository struct {
	Db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{
		Db: db,
	}
}

func (r *NotificationRepository) save(notification *entity.Notification) error {
	_, err := r.Db.Exec("INSERT INTO notifications (id, recipient_id, content, category, read_at, created_at) VALUES ($1, $2, $3, $4, $5, $6)", notification.ID, notification.RecipientID, notification.Content, notification.Category, notification.ReadAt, notification.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
