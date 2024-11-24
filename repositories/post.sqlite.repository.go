package repositories

import (
	"fmt"

	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)


type SqlitePostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *SqlitePostRepository {
	return &SqlitePostRepository{db: db}
}

func (r *SqlitePostRepository) CreatePost(post *models.Post) error {
	fmt.Printf("CreatePost: %+v\n", post)
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(post).Error; err != nil {
			return err
		}

		// Increment post count for subreddit
		if err := tx.Model(&models.Subreddit{}).Where("id = ?", post.SubredditID).Update("post_count", gorm.Expr("post_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *SqlitePostRepository) GetByID(id uint64) (*models.Post, error) {
	var post models.Post
	err := r.db.
		Preload("Subreddit").
		Preload("Author").
		Where("id = ?", id).
		First(&post).Error

	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *SqlitePostRepository) GetPostsBySubreddit(subredditID uint64, limit, offset int32) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.
		Preload("Subreddit").
		Preload("Author").
		Where("subreddit_id = ?", subredditID).
		Order("created_at DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *SqlitePostRepository) GetPostsByUser(userID uint64, limit, offset int32) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.
		Preload("Subreddit").
		Preload("Author").
		Where("author_id = ?", userID).
		Order("created_at DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *SqlitePostRepository) UpdatePostVote(postID uint64, vote int32) error {
	return r.db.Model(&models.Post{}).Where("id = ?", postID).Update("votes", gorm.Expr("votes + ?", vote)).Error
}
