package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetCommentsByPostID(postID string) ([]models.Comment, error)
	GetCommentsByUserID(userID string) ([]models.Comment, error)
	Delete(commentID string) error
}