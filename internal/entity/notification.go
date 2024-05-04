package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID `json:"id"`
	RecipientID uuid.UUID `json:"recipient_id"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	ReadAt      time.Time `json:"read_at"`
	CanceledAt  time.Time `json:"canceled_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewNotification(recipientID uuid.UUID, content, category string) (*Notification, error) {
	notification := &Notification{
		ID:          uuid.New(),
		RecipientID: recipientID,
		Content:     content,
		Category:    category,
		ReadAt:      time.Time{},
		CanceledAt:  time.Time{},
		CreatedAt:   time.Now(),
	}
	if err := notification.validate(); err != nil {
		return nil, err
	}

	return notification, nil
}

func (n *Notification) Read() {
	n.ReadAt = time.Now()
}

func (n *Notification) Unread() {
	n.ReadAt = time.Time{}
}

func (n *Notification) Cancel() {
	n.CanceledAt = time.Now()
}

func (n *Notification) validate() error {
	if n.RecipientID == uuid.Nil {
		return fmt.Errorf("recipientID is required")
	}
	if n.Content == "" {
		return fmt.Errorf("content is required")
	}
	if n.Category == "" {
		return fmt.Errorf("category is required")
	}
	return nil
}
