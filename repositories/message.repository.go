package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type MessageRepository interface {
	SendMessage(text string, fromId uint, toId uint) error
	GetMessages(fromId, toId uint) ([]*models.Message, error)
}
