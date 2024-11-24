package services

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

type MessageService struct {
	msgRepo repositories.MessageRepository
}

func NewMessageService(msgRepo repositories.MessageRepository) *MessageService {
	return &MessageService{
		msgRepo: msgRepo,
	}
}

func (ms *MessageService) SendMessage(text string, fromId, toId uint64) error {
	return ms.msgRepo.SendMessage(text, fromId, toId)
}

func (ms *MessageService) GetMessages(fromId, toId uint64) ([]*models.Message, error) {
	return ms.msgRepo.GetMessages(fromId, toId)
}