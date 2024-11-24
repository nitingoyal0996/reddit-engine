package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type CommentRepository interface {
    Create(comment *models.Comment) error
    GetByID(id uint64) (*models.Comment, error)
    GetCommentsByPost(postID uint64, limit, offset int32) ([]*models.Comment, error)
    GetChildComments(parentID uint64) ([]*models.Comment, error)
    UpdateCommentVote(commentID uint64, vote int32) error
}
