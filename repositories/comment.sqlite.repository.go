package repositories

import (
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
    return r.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(comment).Error; err != nil {
            return err
        }

        // Increment comment count for post
        if err := tx.Model(&models.Post{}).Where("id = ?", comment.PostID).
            Update("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
            return err
        }

        return nil
    })
}

func (r *SqliteCommentRepository) GetByID(id uint64) (*models.Comment, error) {
    var comment models.Comment
    err := r.db.
        Preload("User").
        Preload("Post").
        Preload("Children").
        Where("id = ?", id).
        First(&comment).Error

    if err != nil {
        return nil, err
    }
    return &comment, nil
}

func (r *SqliteCommentRepository) GetCommentsByPost(postID uint64, limit, offset int32) ([]*models.Comment, error) {
    var comments []*models.Comment
    err := r.db.
        Preload("User").
        Preload("Children").
        Where("post_id = ? AND parent_id IS NULL", postID).
        Order("created_at DESC").
        Limit(int(limit)).
        Offset(int(offset)).
        Find(&comments).Error

    if err != nil {
        return nil, err
    }
    return comments, nil
}

func (r *SqliteCommentRepository) GetChildComments(parentID uint64) ([]*models.Comment, error) {
    var comments []*models.Comment
    err := r.db.
        Preload("User").
        Preload("Children").
        Where("parent_id = ?", parentID).
        Order("created_at DESC").
        Find(&comments).Error

    if err != nil {
        return nil, err
    }
    return comments, nil
}

func (r *SqliteCommentRepository) UpdateCommentVote(commentID uint64, vote int32) error {
        return r.db.Model(&models.Comment{}).Where("id = ?", commentID).Update("votes", gorm.Expr("votes + ?", vote)).Error
}