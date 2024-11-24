package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type SubredditRepository interface {
	CreateSubreddit(sub *models.Subreddit) error
	GetByID(id uint64) (*models.Subreddit, error)
	Subscribe(userID, subredditID uint64) error
	Unsubscribe(userID, subredditID uint64) error
	GetUserSubscriptions(userID uint64) ([]*models.Subreddit, error)
	SearchSubreddits(query string, limit int) ([]*models.Subreddit, error)
}