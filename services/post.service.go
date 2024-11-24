package services

import (
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/models"
)

type PostService struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (ps *PostService) CreatePost(post *models.Post) error {
	return ps.postRepo.Create(post)
}

func (ps *PostService) DeletePost(postId uint) error {
	return ps.postRepo.Delete(postId)
}

