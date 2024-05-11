package repository

import (
	"database/sql"
	"github/eovinicius/notification/internal/entity"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	Save(notification *entity.Notification) error
	FindByID(id uuid.UUID) (*entity.Notification, error)
	Update(notification *entity.Notification) error
	FindManyByRecipientID(recipientID uuid.UUID) ([]*entity.Notification, error)
}

type notificationRepository struct {
	Db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *notificationRepository {
	return &notificationRepository{
		Db: db,
	}
}

func (r *notificationRepository) Save(notification *entity.Notification) error {
	_, err := r.Db.Exec("INSERT INTO notifications (id, recipient_id, content, category, read_at, created_at) VALUES ($1, $2, $3, $4, $5, $6)", notification.ID, notification.RecipientID, notification.Content, notification.Category, notification.ReadAt, notification.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *notificationRepository) FindByID(id uuid.UUID) (*entity.Notification, error) {
	var notification entity.Notification
	err := r.Db.QueryRow("SELECT id, recipient_id, content, category, read_at, created_at FROM notifications WHERE id = $1", id).Scan(&notification.ID, &notification.RecipientID, &notification.Content, &notification.Category, &notification.ReadAt, &notification.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) Update(notification *entity.Notification) error {
	_, err := r.Db.Exec("UPDATE notifications SET recipient_id = $1, content = $2, category = $3, read_at = $4, canceled_at = $5, created_at = $6 WHERE id = $7", notification.RecipientID, notification.Content, notification.Category, notification.ReadAt, notification.CanceledAt, notification.CreatedAt, notification.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *notificationRepository) FindManyByRecipientID(recipientID string) ([]*entity.Notification, error) {
	rows, err := r.Db.Query("SELECT id, recipient_id, content, category, read_at, created_at FROM notifications WHERE recipient_id = $1", recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []*entity.Notification
	for rows.Next() {
		var notification entity.Notification
		err := rows.Scan(&notification.ID, &notification.RecipientID, &notification.Content, &notification.Category, &notification.ReadAt, &notification.CreatedAt)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	return notifications, nil
}
