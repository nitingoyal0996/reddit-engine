package repositories

import (
	"errors"
	"fmt"

	"github.com/nitingoyal0996/reddit-clone/models"
	"gorm.io/gorm"
)

type SqliteSubredditRepository struct {
    db *gorm.DB
}

func NewSubredditRepository(db *gorm.DB) *SqliteSubredditRepository {
    return &SqliteSubredditRepository{db: db}
}

// Create a new subreddit
func (r *SqliteSubredditRepository) CreateSubreddit(subreddit *models.Subreddit) error {
    // PRINT INCOMING DATA WITH FIELDS
    fmt.Printf("CreateSubreddit: %+v\n", subreddit)
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Create the subreddit
        if err := tx.Create(subreddit).Error; err != nil {
            return err
        }

        // Automatically subscribe the creator
        subscription := models.UserSubredditSubscription{
            UserID:      subreddit.CreatorID,
            SubredditID: subreddit.ID,
        }
        
        return tx.Create(&subscription).Error
    })
}

// Get subreddit by ID
func (r *SqliteSubredditRepository) GetByID(id uint64) (*models.Subreddit, error) {
	var subreddit models.Subreddit
	err := r.db.
		Preload("Creator").
		Preload("Rules", func(db *gorm.DB) *gorm.DB {
			return db.Order("order_number ASC")
		}).
		Where("id = ? AND is_active = ?", id, true).
		First(&subreddit).Error
	
	if err != nil {
		return nil, err
	}
	return &subreddit, nil
}

// Subscribe user to subreddit
func (r *SqliteSubredditRepository) Subscribe(userID, subredditID uint64) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        // Check if subscription already exists
        var exists int64
        tx.Model(&models.UserSubredditSubscription{}).
            Where("user_id = ? AND subreddit_id = ?", userID, subredditID).
            Count(&exists)
            
        if exists > 0 {
            return errors.New("already subscribed")
        }

        // Create subscription
        subscription := models.UserSubredditSubscription{
            UserID:      userID,
            SubredditID: subredditID,
        }
        
        if err := tx.Create(&subscription).Error; err != nil {
            return err
        }

        // Update subscriber count
        return tx.Model(&models.Subreddit{}).
            Where("id = ?", subredditID).
            UpdateColumn("subscriber_count", gorm.Expr("subscriber_count + ?", 1)).
            Error
    })
}

// Unsubscribe user from subreddit
func (r *SqliteSubredditRepository) Unsubscribe(userID, subredditID uint64) error {
    return r.db.Transaction(func(tx *gorm.DB) error {
        result := tx.Where("user_id = ? AND subreddit_id = ?", userID, subredditID).
            Delete(&models.UserSubredditSubscription{})
            
        if result.RowsAffected == 0 {
            return errors.New("subscription not found")
        }

        // Update subscriber count
        return tx.Model(&models.Subreddit{}).
            Where("id = ?", subredditID).
            UpdateColumn("subscriber_count", gorm.Expr("subscriber_count - ?", 1)).
            Error
    })
}

// Get user's subscribed subreddits
func (r *SqliteSubredditRepository) GetUserSubscriptions(userID uint64) ([]*models.Subreddit, error) {
    var subreddits []*models.Subreddit
    err := r.db.
        Joins("JOIN user_subreddit_subscriptions ON subreddits.id = user_subreddit_subscriptions.subreddit_id").
        Where("user_subreddit_subscriptions.user_id = ? AND subreddits.is_active = ?", userID, true).
        Preload("Creator").
        Find(&subreddits).Error
    return subreddits, err
}

// Search subreddits
func (r *SqliteSubredditRepository) SearchSubreddits(query string, limit int) ([]*models.Subreddit, error) {
    var subreddits []*models.Subreddit
    err := r.db.
        Where("name LIKE ? AND is_active = ?", "%"+query+"%", true).
        Or("description LIKE ? AND is_active = ?", "%"+query+"%", true).
        Limit(limit).
        Preload("Creator").
        Find(&subreddits).Error
    return subreddits, err
}
