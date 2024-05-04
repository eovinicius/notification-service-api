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

func (r *NotificationRepository) Save(notification *entity.Notification) error {
	_, err := r.Db.Exec("INSERT INTO notifications (id, recipient_id, content, category, read_at, created_at) VALUES ($1, $2, $3, $4, $5, $6)", notification.ID, notification.RecipientID, notification.Content, notification.Category, notification.ReadAt, notification.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *NotificationRepository) FindByID(id string) (*entity.Notification, error) {
	var notification entity.Notification
	err := r.Db.QueryRow("SELECT id, recipient_id, content, category, read_at, created_at FROM notifications WHERE id = $1", id).Scan(&notification.ID, &notification.RecipientID, &notification.Content, &notification.Category, &notification.ReadAt, &notification.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *NotificationRepository) Update(notification *entity.Notification) error {
	_, err := r.Db.Exec("UPDATE notifications SET recipient_id = $1, content = $2, category = $3, read_at = $4, canceled_at = $5, created_at = $6 WHERE id = $7", notification.RecipientID, notification.Content, notification.Category, notification.ReadAt, notification.CanceledAt, notification.CreatedAt, notification.ID)
	if err != nil {
		return err
	}

	return nil
}
