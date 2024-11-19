package actors

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
)

type AuthActor struct {
    service     *services.AuthService
}

func NewAuthActor(userRepo *repositories.SqliteUserRepository, jwtSecret string) *AuthActor {
    return &AuthActor {
        service: services.NewAuthService(userRepo, jwtSecret),
    }
}

func (auth *AuthActor) Receive(context actor.Context) {
	switch actorMsg := context.Message().(type) {
	case *messages.RegisterRequest: 
		user, err := auth.service.RegisterNewUser(actorMsg.Username, actorMsg.Email, actorMsg.Password)
        if err != nil {
            context.Respond(&messages.RegisterResponse{Error: err.Error()})
            return
        }
        context.Respond(&messages.RegisterResponse{ID: user.ID, Username: user.Username})
    case *messages.LoginRequest:
        token, err := auth.service.Login(actorMsg.Username, actorMsg.Password)
        if err != nil {
            context.Respond(&messages.LoginResponse{Error: err.Error()})
            return
        }
        context.Respond(&messages.LoginResponse{Token: token})
    case *messages.TokenValidationRequest:
        claims, err := auth.service.ValidateToken(actorMsg.Token)
        if err != nil {
            context.Respond(&messages.TokenValidationResponse{Error: err.Error()})
            return
        }
        context.Respond(&messages.TokenValidationResponse{Valid: true, Claims: claims})
	}

}

// Register
// ValidateToken
// Login > GenerateToken