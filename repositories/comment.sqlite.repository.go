package repositories

import (
	"errors"
	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqliteCommentRepository struct {
	db *gorm.DB
}	

func NewCommentRepository(db *gorm.DB) *SqliteCommentRepository {
	return &SqliteCommentRepository{db: db}
}

func (r *SqliteCommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *SqliteCommentRepository) GetCommentsByPostID(postID string) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *SqliteCommentRepository) GetCommentById(commentID string) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(&comment, commentID).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *SqliteCommentRepository) DeleteComment(commentID string) error {
	var comment models.Comment
	if err := r.db.First(&comment, commentID).Error; err != nil {
		return err
	}
	if comment.IsDeleted {
		return errors.New("comment already deleted")
	}
	comment.IsDeleted = true
	return r.db.Save(&comment).Error
}

