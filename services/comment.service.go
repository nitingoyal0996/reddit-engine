package services

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

type CommentService struct {
    commentRepo repositories.CommentRepository
}

type CreateCommentInput struct {
    Content  string
    UserID   uint64
    PostID   uint64
    ParentID *uint64
}

func NewCommentService(repo repositories.CommentRepository) *CommentService {
    return &CommentService{
		commentRepo: repo,
	}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
    return s.commentRepo.Create(comment)
}

func (s *CommentService) GetCommentThread(postID uint64) ([]*models.Comment, error) {
    // Get root level comments with a reasonable limit
    comments, err := s.commentRepo.GetCommentsByPost(postID, 100, 0)
    if err != nil {
        return nil, err
    }
    
    // Load child comments recursively
    for _, comment := range comments {
        if err := s.loadChildComments(comment); err != nil {
            return nil, err
        }
    }
    
    return comments, nil
}

func (s *CommentService) loadChildComments(comment *models.Comment) error {
    children, err := s.commentRepo.GetChildComments(uint64(*comment.ParentID))
    if err != nil {
        return err
    }
    
    comment.Children = make([]models.Comment, len(children))
    for i, child := range children {
        comment.Children[i] = *child
    }
    for _, child := range children {
        if err := s.loadChildComments(child); err != nil {
            return err
        }
    }
    return nil
}

func (s *CommentService) GetCommentByID(id uint64) (*models.Comment, error) {
	return s.commentRepo.GetByID(id)
}

func (s *CommentService) UpdateCommentVote(commentID uint64, vote int32) error {
	return s.commentRepo.UpdateCommentVote(commentID, vote)
}

func (s *CommentService) GetCommentsByPost(postID uint64, limit, offset int32) ([]*models.Comment, error) {
	return s.commentRepo.GetCommentsByPost(postID, limit, offset)
}