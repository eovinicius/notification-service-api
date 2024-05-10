package dto

import "github.com/google/uuid"

type SendNotificationRequest struct {
	RecipientID uuid.UUID `json:"recipient_id" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Category    string    `json:"category" binding:"required"`
}
