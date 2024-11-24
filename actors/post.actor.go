package actors

import (
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type PostActor struct {
	postService *services.PostService
}	

func NewPostActor(postRepo *repositories.SqlitePostRepository) *PostActor {
	return &PostActor{
		postService: services.NewPostService(postRepo),
	}
}