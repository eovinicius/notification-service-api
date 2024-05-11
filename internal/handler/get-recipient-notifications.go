package handler

import (
	"github/eovinicius/notification/internal/entity"
	usecases "github/eovinicius/notification/internal/use-cases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetRecipientNotifications struct {
	getRecipientNotifications *usecases.GetRecipientNotifications
}

func NewGetRecipientNotifications(getRecipientNotifications *usecases.GetRecipientNotifications) *GetRecipientNotifications {
	return &GetRecipientNotifications{getRecipientNotifications: getRecipientNotifications}
}

func (h *GetRecipientNotifications) Handle(ctx *gin.Context) ([]*entity.Notification, error) {

	recipientID := ctx.Param("recipientID")

	uuid, err := uuid.Parse(recipientID)

	notifications, err := h.getRecipientNotifications.Execute(uuid)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
