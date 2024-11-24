package actors

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/proto"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

func convertUint64ToUint(value *uint64) *uint {
	if value == nil {
		return nil
	}
	convertedValue := uint(*value)
	return &convertedValue
}


type CommentActor struct {
	commentService *services.CommentService
}

func NewCommentActor(commentRepo *repositories.SqliteCommentRepository) *CommentActor {
	return &CommentActor{
		commentService: services.NewCommentService(commentRepo),
	}
}

func (comment *CommentActor) Receive(context actor.Context) {
	msg := context.Message()
	fmt.Printf("CommentActor Received message %v\n", msg)
	switch actorMsg := msg.(type) {
	case *proto.CreateCommentRequest:
		comment.CreateComment(context, actorMsg)
	case *proto.GetCommentRequest:
		comment.GetComment(context, actorMsg)
	// case *proto.GetCommentsByPostRequest:
	// 	comment.GetCommentsByPost(context, actorMsg)
	// case *proto.GetPostByUserRequest:
	// 	comment.GetCommentsByUser(context, actorMsg)
	// case *proto.UpdateCommentVoteRequest:
	// 	comment.UpdateCommentVote(context, actorMsg)
	default:
		println("Unknown message to PostActor")
	}
}

func (actor *CommentActor) CreateComment(context actor.Context, actorMsg *proto.CreateCommentRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreateCommentResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreateCommentResponse{Error: "Invalid token"})
	} else {
		newComment := &models.Comment{
			Content: actorMsg.Content,
			PostID: actorMsg.PostId,
			UserID: actorMsg.UserId,
			ParentID: convertUint64ToUint(actorMsg.ParentId),
		}
		if err := actor.commentService.CreateComment(newComment); err != nil {
			context.Respond(&proto.CreateCommentResponse{Error: err.Error()})
		} else {
			context.Respond(&proto.CreateCommentResponse{Error: ""})
		}
	}
}

func (actor *CommentActor) GetComment(context actor.Context, actorMsg *proto.GetCommentRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreateCommentResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreateCommentResponse{Error: "Invalid token"})
	} else {
		retrievedComment, err := actor.commentService.GetCommentByID(actorMsg.CommentId)
		if err != nil {
			context.Respond(&proto.GetCommentResponse{Comment: nil, Error: err.Error()})
		} else {
			protoComment := &proto.Comment{
				Id:      retrievedComment.ID,
				Content: retrievedComment.Content,
				PostId:  retrievedComment.PostID,
				UserId:  retrievedComment.UserID,
				ParentId: func() *uint64 {
					if retrievedComment.ParentID != nil {
						parentID := uint64(*retrievedComment.ParentID)
						return &parentID
					}
					return nil
				}(),
			}
			context.Respond(&proto.GetCommentResponse{Error: "", Comment: protoComment})
		}
	}
}