package repositories

import (
	"errors"
	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqlitePostRepository struct {
	db *gorm.DB
}	

func NewPostRepository(db *gorm.DB) *SqlitePostRepository {
	return &SqlitePostRepository{db: db}
}

func (r *SqlitePostRepository) Create(post *models.Post) error {
	var existingPost models.Post
	if r.db.Where("title = ?", post.Title).First(&existingPost); existingPost.ID != 0 {
		return errors.New("post already exists")
	}

	return r.db.Create(post).Error
}

func (r *SqlitePostRepository) GetPostById(PostId uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, PostId).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *SqlitePostRepository) Delete(postId uint) error {
	return r.db.Delete(&models.Post{ID: postId}).Error
}