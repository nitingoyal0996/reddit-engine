package actors

import (
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type KarmaActor struct {
	karmaService 	*services.KarmaService
	authPID	 		*actor.PID
	rootContext 	*actor.RootContext
}

func NewKarmaActor(userRepo *repositories.SqliteUserRepository, authPID *actor.PID, rootContext *actor.RootContext) *KarmaActor {
	return &KarmaActor{
		karmaService: services.NewKarmaService(userRepo),
		authPID: authPID,
		rootContext: rootContext,
	}
}

func (karma *KarmaActor) Receive(context actor.Context) {
	switch actorMsg := context.Message().(type) {
	case *messages.KarmaRequest:
		// validate the token
		future := karma.rootContext.RequestFuture(karma.authPID, &messages.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
		result, err := future.Result()
		if err != nil {
			context.Respond(&messages.KarmaResponse{Error: err.Error()})
			return
		}
		validationResponse, ok := result.(*messages.TokenValidationResponse)
		if !ok {
			context.Respond(&messages.KarmaResponse{Error: "Invalid response from auth actor"})
			return
		}
		if !validationResponse.Valid {
			context.Respond(&messages.KarmaResponse{Error: "Invalid token"})
			return
		}
		// update karma
		if err := karma.karmaService.UpdateKarma(validationResponse.Claims.UserId, actorMsg.Amount); err != nil {
			context.Respond(&messages.KarmaResponse{Error: err.Error()})
			return
		}
		context.Respond(&messages.KarmaResponse{Error: ""})
	}
}