package repositories

import "github.com/nitingoyal0996/reddit-clone/models"

type SubRepository interface {
	Create(sub *models.Subreddit) error
	GetSubredditByName(SubredditName string) (*models.Subreddit, error)
	GetSubredditById(SubredditId uint) (*models.Subreddit, error)
}