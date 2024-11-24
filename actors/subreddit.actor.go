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

type SubredditActor struct {
	subredditService *services.SubredditService
}

func NewSubredditActor(subRepo *repositories.SqliteSubredditRepository) *SubredditActor {
	return &SubredditActor{
		subredditService: services.NewSubredditService(subRepo),
	}
}

func (subreddit * SubredditActor) Receive (context actor.Context) {
	msg := context.Message()
	fmt.Printf("Subreddit actor received message: %T\n", msg)
	switch actorMsg := msg.(type) {
	case *actor.Started:
		println("SubredditActor started")
	case *actor.Stopping:
		println("SubredditActor stopping")
	case *actor.Terminated:
		println("SubredditActor terminated")
	case *proto.CreateSubredditRequest:
		subreddit.CreateSubreddit(context, actorMsg)
	case *proto.SubscriptionRequest:
		println("Subscription request received")
		subreddit.Subscribe(context, actorMsg)
	case *proto.UnsubscribeRequest:
		println("Unsubscription request received")
		subreddit.Unsubscribe(context, actorMsg)
	default:
		println("Unknown message to SubredditActor")
	}
}

func (subreddit *SubredditActor) CreateSubreddit(context actor.Context, actorMsg *proto.CreateSubredditRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.CreateSubredditResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.CreateSubredditResponse{Error: "Invalid token"})
	} else {
		fmt.Println("Token validated successfully")
		// create subreddit
		if err := subreddit.subredditService.CreateSubreddit(&models.Subreddit{Name: actorMsg.Name, Description: actorMsg.Description, CreatorID: actorMsg.CreatorId}); err != nil {
			context.Respond(&proto.CreateSubredditResponse{Error: err.Error()})
		}
		context.Respond(&proto.CreateSubredditResponse{Error: ""})
	}
}

func (subreddit *SubredditActor) Subscribe(context actor.Context, actorMsg *proto.SubscriptionRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.SubscriptionResponse{Success: false, Message: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.SubscriptionResponse{Success: false, Message: "Invalid token"})
	} else {
		fmt.Println("Token validated successfully")
		// subscribe to subreddit
		if err := subreddit.subredditService.Subscribe(validationResponse.Claims.UserId, actorMsg.SubredditId); err != nil {
			context.Respond(&proto.SubscriptionResponse{Success: false, Message: err.Error()})
		}
		context.Respond(&proto.SubscriptionResponse{Success: true, Message: ""})
	}
}

func (subreddit *SubredditActor) Unsubscribe(context actor.Context, actorMsg *proto.UnsubscribeRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.SubscriptionResponse{Success: false, Message: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.SubscriptionResponse{Success: false, Message: "Invalid token"})
	} else {
		fmt.Println("Token validated successfully")
		// unsubscribe to subreddit
		if err := subreddit.subredditService.Unsubscribe(validationResponse.Claims.UserId, actorMsg.SubredditId); err != nil {
			context.Respond(&proto.SubscriptionResponse{Success: false, Message: err.Error()})
		}
		context.Respond(&proto.SubscriptionResponse{Success: true, Message: ""})
	}
}