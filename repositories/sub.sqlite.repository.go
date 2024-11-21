package repositories

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqltieSubRepository struct {
	db *gorm.DB
}

func NewSqliteSubRepository(db *gorm.DB) *SqltieSubRepository {
	return &SqltieSubRepository{db: db}
}

func (r *SqltieSubRepository) Create(sub *models.Subreddit) error {
	return r.db.Create(sub).Error
}