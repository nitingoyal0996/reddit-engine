package actors

import (
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/models"
	"github.com/nitingoyal0996/reddit-clone/proto"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostActor struct {
	postService *services.PostService
}

func NewPostActor(postRepo *repositories.SqlitePostRepository) *PostActor {
	return &PostActor{
		postService: services.NewPostService(postRepo),
	}
}

func (post *PostActor) Receive(context actor.Context) {
	msg := context.Message()
	switch actorMsg := msg.(type) {
	case *proto.CreatePostRequest:
		post.CreatePost(context, actorMsg)
	case *proto.GetPostRequest:
		post.GetPost(context, actorMsg)
	case *proto.GetPostsBySubredditRequest:
		post.GetPostsBySubreddit(context, actorMsg)
	case *proto.GetPostByUserRequest:
		post.GetPostsByUser(context, actorMsg)
	case *proto.UpdatePostVoteRequest:
		post.UpdatePostVote(context, actorMsg)
	default:
		println("Unknown message to PostActor")
	}
}

func (post *PostActor) CreatePost(context actor.Context, actorMsg *proto.CreatePostRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreatePostResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreatePostResponse{Error: "Invalid token"})
	} else {
		newPost := &models.Post{
			Title:       actorMsg.Title,
			Content:     actorMsg.Content,
			SubredditID: actorMsg.SubredditId,
			AuthorID:    actorMsg.AuthorId,
		}
		err := post.postService.CreatePost(newPost)
		if err != nil {
			context.Respond(&proto.CreatePostResponse{Error: err.Error()})
		}
		context.Respond(&proto.CreatePostResponse{Error: ""})
	}
}

func (post *PostActor) GetPost(context actor.Context, actorMsg *proto.GetPostRequest) {
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreatePostResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreatePostResponse{Error: "Invalid token"})
	} else {
		retrievedPost, err := post.postService.GetPostByID(actorMsg.PostId)
		if err != nil {
			context.Respond(&proto.GetPostResponse{})
			return
		}
		protoPost := &proto.Post{
			Id:          retrievedPost.ID,
			Title:       retrievedPost.Title,
			Content:     retrievedPost.Content,
			SubredditId: retrievedPost.SubredditID,
			AuthorId:    retrievedPost.AuthorID,
			CreatedAt:   timestamppb.New(retrievedPost.CreatedAt),
		}
		context.Respond(&proto.GetPostResponse{Post: protoPost})
	}
}

func (post *PostActor) GetPostsBySubreddit(context actor.Context, actorMsg *proto.GetPostsBySubredditRequest) {
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreatePostResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreatePostResponse{Error: "Invalid token"})
	} else {
		posts, err := post.postService.GetPostsBySubreddit(actorMsg.SubredditId, actorMsg.Limit, actorMsg.Offset)
		if err != nil {
			context.Respond(&proto.GetPostsBySubredditResponse{Posts: nil})
		}
		var protoPosts []*proto.Post
		for _, post := range posts {
			protoPosts = append(protoPosts, &proto.Post{
				Id:          post.ID,
				Title:       post.Title,
				Content:     post.Content,
				SubredditId: post.SubredditID,
				AuthorId:    post.AuthorID,
				CreatedAt:   timestamppb.New(post.CreatedAt),
			})
		}
		context.Respond(&proto.GetPostsBySubredditResponse{Posts: protoPosts})
	}
}

func (post *PostActor) GetPostsByUser(context actor.Context, actorMsg *proto.GetPostByUserRequest) {
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreatePostResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreatePostResponse{Error: "Invalid token"})
	} else {
		posts, err := post.postService.GetPostsByUser(validationResponse.Claims.UserId, actorMsg.Limit, actorMsg.Offset)
		if err != nil {
			context.Respond(&proto.GetPostByUserResponse{Posts: nil})
		}
		var protoPosts []*proto.Post
		for _, post := range posts {
			protoPosts = append(protoPosts, &proto.Post{
				Id:          post.ID,
				Title:       post.Title,
				Content:     post.Content,
				SubredditId: post.SubredditID,
				AuthorId:    post.AuthorID,
				CreatedAt:   timestamppb.New(post.CreatedAt),
			})
		}
		context.Respond(&proto.GetPostByUserResponse{Posts: protoPosts})
	}
}

func (post *PostActor) UpdatePostVote(context actor.Context, actorMsg *proto.UpdatePostVoteRequest) {
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreatePostResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreatePostResponse{Error: "Invalid token"})
	} else {
		voteCount := -1
		if actorMsg.Upvote { voteCount = 1  }
		err := post.postService.UpdatePostVote(actorMsg.PostId, int32(voteCount))
		if err != nil {
			context.Respond(&proto.UpdatePostVoteResponse{Error: err.Error()})
		} else {
			// fetch post author by get by post id
			post, err := post.postService.GetPostByID(actorMsg.PostId)
			if err != nil {
				context.Respond(&proto.UpdatePostVoteResponse{Error: err.Error()})
			}
			karmaActor := cluster.GetCluster(context.ActorSystem()).Get("karma", "Karma")
			karmaFuture := context.RequestFuture(karmaActor, &proto.KarmaRequest{
				Token: actorMsg.Token,
				UserId: post.AuthorID,
				Amount: int32(voteCount),
			}, 5*time.Second)
			_, err = karmaFuture.Result()
			if err != nil {
				context.Respond(&proto.UpdatePostVoteResponse{Error: err.Error()})
			}	
		}
	}
	context.Respond(&proto.UpdatePostVoteResponse{Error: ""})
}
