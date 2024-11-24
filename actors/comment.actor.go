package actors

import (
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)	

type CommentActor struct {
	commentService *services.CommentService
}	

func NewCommentActor(repo repositories.CommentRepository) *CommentActor {
	return &CommentActor{
		commentService: services.NewCommentService(repo),
	}
}