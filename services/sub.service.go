package services

import (
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

type SubredditService struct {
	subredditRepo 	repositories.SubredditRepository
}

func NewSubredditService(subredditRepo repositories.SubredditRepository) *SubredditService {
	return &SubredditService{
		subredditRepo: subredditRepo,
	}
}

func (s *SubredditService) CreateSubreddit(reddit *models.Subreddit) (subredditID uint64, string error) {
	return s.subredditRepo.CreateSubreddit(reddit)
}

func (s *SubredditService) GetSubredditByID(id uint64) (*models.Subreddit, error) {
	return s.subredditRepo.GetByID(id)
}

func (s *SubredditService) Subscribe(userID, subredditID uint64) error {
	return s.subredditRepo.Subscribe(userID, subredditID)
}

func (s *SubredditService) Unsubscribe(userID, subredditID uint64) error {
	return s.subredditRepo.Unsubscribe(userID, subredditID)
}

func (s *SubredditService) GetSubscribedSubreddits(userID uint64) ([]*models.Subreddit, error) {
	return s.subredditRepo.GetUserSubscriptions(userID)
}

func (s *SubredditService) SearchSubreddits(query string, limit int) ([]*models.Subreddit, error) {
	return s.subredditRepo.SearchSubreddits(query, limit)
}