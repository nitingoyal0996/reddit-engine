package repositories

import (
	"github.com/nitingoyal0996/reddit-clone/models"
)

type PostRepository interface {
	GetPostById(PostId uint) (*models.Post, error)
	Create(post *models.Post) error
	Delete(postId uint) error
}