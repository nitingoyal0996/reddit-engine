package services

import (
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/models"
)

type CommentService struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) *CommentService {
	return &CommentService{
		commentRepo: commentRepo,
	}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.commentRepo.Create(comment)
}

func (s *CommentService) GetCommentsByPostID(postID string) ([]models.Comment, error) {
	return s.commentRepo.GetCommentsByPostID(postID)
}

func (s *CommentService) GetCommentsByUserID(userID string) ([]models.Comment, error) {
	return s.commentRepo.GetCommentsByUserID(userID)
}

func (s *CommentService) Delete(commentID string) error {
	return s.commentRepo.Delete(commentID)
}