package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type MessageRepository interface {
	SendMessage(text string, fromId, toId uint64) error
	GetMessages(fromId, toId uint64) ([]*models.Message, error)
}
