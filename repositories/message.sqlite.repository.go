package repositories

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqliteMessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *SqliteMessageRepository {
	return &SqliteMessageRepository{db: db}
}

func (r *SqliteMessageRepository) SendMessage(text string, fromId, toId uint64) error {
	message := &models.Message{
		Text:   text,
		FromId: fromId,
		ToId:   toId,
	}
	return r.db.Create(message).Error
}

func (r *SqliteMessageRepository) GetMessages(fromId, toId uint64) ([]*models.Message, error) {
	var messages []*models.Message
	if err := r.db.Where("from_id = ? AND to_id = ? OR from_id = ? AND to_id = ?", fromId, toId, toId, fromId).Order("created_at desc").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
