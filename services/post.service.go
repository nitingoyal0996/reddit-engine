package services

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

type PostService struct {
	postRepo 	repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.postRepo.CreatePost(post)
}

func (s *PostService) GetPostByID(id uint64) (*models.Post, error) {
	return s.postRepo.GetByID(id)
}

func (s *PostService) GetPostsBySubreddit(subredditID uint64, limit, offset int32) ([]*models.Post, error) {
	return s.postRepo.GetPostsBySubreddit(subredditID, limit, offset)
}

func (s *PostService) GetPostsByUser(userID uint64, limit, offset int32) ([]*models.Post, error) {
	return s.postRepo.GetPostsByUser(userID, limit, offset)
}

func (s *PostService) UpdatePostVote(postID uint64, vote int32) error {
	return s.postRepo.UpdatePostVote(postID, vote)
}