package actors

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type KarmaActor struct {
	karmaService 	*services.KarmaService
}

func NewKarmaActor(userRepo *repositories.SqliteUserRepository) *KarmaActor {
	return &KarmaActor{
		karmaService: services.NewKarmaService(userRepo),
	}
}

func (karma *KarmaActor) Receive(context actor.Context) {
	switch actorMsg := context.Message().(type) {
	case *messages.KarmaRequest:
		// update karma
		if err := karma.karmaService.UpdateKarma(actorMsg.UserId, actorMsg.Amount); err != nil {
			context.Respond(&messages.KarmaResponse{Error: err.Error()})
			return
		}
		context.Respond(&messages.KarmaResponse{Error: ""})
	}
}