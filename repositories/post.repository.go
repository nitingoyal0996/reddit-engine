package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type PostRepository interface {
	CreatePost(post *models.Post) error
	GetByID(id uint64) (*models.Post, error)
	GetPostsBySubreddit(subredditID uint64, limit, offset int32) ([]*models.Post, error)
	GetPostsByUser(userID uint64, limit, offset int32) ([]*models.Post, error)
	UpdatePostVote(postID uint64, vote int32) error
}