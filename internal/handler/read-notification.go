package handler

import (
	usecases "github/eovinicius/notification/internal/use-cases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReadNotification struct {
	readNotification *usecases.ReadNotification
}

func NewReadNotification(readNotification *usecases.ReadNotification) *ReadNotification {
	return &ReadNotification{readNotification: readNotification}
}

func (h *ReadNotification) Handle(ctx *gin.Context) {

	id := ctx.Param("notificationID")

	notificatonId, err := uuid.Parse(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err = h.readNotification.Execute(notificatonId, uuid.New()); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Notification marked as read"})
}
