package actors

import (
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type SubredditActor struct {
	subredditService *services.SubredditService
}

func NewSubredditActor(subredditRepo *repositories.SqliteSubRepository) *SubredditActor {
	return &SubredditActor{
		subredditService: services.NewSubredditService(subredditRepo),
	}
}