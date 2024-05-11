package handler

import (
	usecases "github/eovinicius/notification/internal/use-cases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CancelNotification struct {
	cancelNotification *usecases.CancelNotification
}

func NewCancelNotification(cancelNotification *usecases.CancelNotification) *CancelNotification {
	return &CancelNotification{cancelNotification: cancelNotification}
}

func (h *CancelNotification) Handle(ctx *gin.Context) {

	recipientID := ctx.Param("notificationID")

	uuid, err := uuid.Parse(recipientID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err = h.cancelNotification.Execute(uuid); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Notification canceled"})
}
