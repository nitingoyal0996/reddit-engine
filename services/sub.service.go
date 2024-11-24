package services

import(
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/repositories"
) 

type SubredditService struct {
	subRepo repositories.SubRepository
}

func NewSubredditService(subRepo repositories.SubRepository) *SubredditService {
	return &SubredditService{
		subRepo: subRepo,
	}
}

func (s *SubredditService) CreateSubreddit(sub *models.Subreddit) error {
	return s.subRepo.Create(sub)
}