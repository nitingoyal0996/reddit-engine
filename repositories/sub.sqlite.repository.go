package repositories

import (
	"errors"
	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqliteSubRepository struct {
	db *gorm.DB
}

func NewSubRepository(db *gorm.DB) *SqliteSubRepository {
	return &SqliteSubRepository{db: db}
}

func (r *SqliteSubRepository) GetSubredditByName (SubredditName string) (*models.Subreddit, error) {
	var subreddit models.Subreddit
	if err := r.db.Where("name = ?", SubredditName).First(&subreddit).Error; err != nil {
		return nil, err
	}

	return &subreddit, nil
}

func (r *SqliteSubRepository) GetSubredditById (SubredditId uint) (*models.Subreddit, error) {
	var subreddit models.Subreddit
	if err := r.db.First(&subreddit, SubredditId).Error; err != nil {
		return nil, err
	}

	return &subreddit, nil
}

func (r *SqliteSubRepository) Create(sub *models.Subreddit) error {

	var existingSubreddit models.Subreddit
	if r.db.Where("name = ?", sub.Name).First(&existingSubreddit); existingSubreddit.ID != 0 {
		return errors.New("subreddit already exists")
	}

	return r.db.Create(sub).Error
}