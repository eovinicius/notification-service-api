package handler

import (
	"github/eovinicius/notification/internal/dto"
	usecases "github/eovinicius/notification/internal/use-cases"

	"github.com/gin-gonic/gin"
)

type SendNotification struct {
	sendNotification *usecases.SendNotification
}

func NewSendNotification(sendNotification *usecases.SendNotification) *SendNotification {
	return &SendNotification{sendNotification: sendNotification}
}

func (h *SendNotification) Handle(ctx *gin.Context) {
	var request dto.SendNotificationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.sendNotification.Execute(request.RecipientID, request.Content, request.Category)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Notification send successfully"})
}
