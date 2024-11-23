package actors

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type KarmaActor struct {
	karmaService *services.KarmaService
}

func NewKarmaActor(userRepo *repositories.SqliteUserRepository) *KarmaActor {
	return &KarmaActor{
		karmaService: services.NewKarmaService(userRepo),
	}
}

func (karma *KarmaActor) Receive(context actor.Context) {
	msg := context.Message()
	fmt.Printf("Received message: %T\n", msg)
	switch actorMsg := msg.(type) {
	case *actor.Started:
		println("KarmaActor started")
	case *actor.Stopping:
		println("KarmaActor stopping")
	case *actor.Terminated:
		println("KarmaActor terminated")
	case *proto.KarmaRequest:
		println("Karma request received")
		karma.UpdateKarma(context, actorMsg)
	default:
		println("Unknown message to KarmaActor")
	}
}

func (karma *KarmaActor) UpdateKarma(context actor.Context, actorMsg *proto.KarmaRequest) {
	// validate request with auth actor
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.KarmaResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.KarmaResponse{Error: "Invalid token"})
	} else {
		fmt.Println("Token validated successfully")
		// update karma
		if err := karma.karmaService.UpdateKarma(uint(validationResponse.Claims.UserId), int(actorMsg.Amount)); err != nil {
			context.Respond(&proto.KarmaResponse{Error: err.Error()})
		}
		context.Respond(&proto.KarmaResponse{Error: ""})
	}
}