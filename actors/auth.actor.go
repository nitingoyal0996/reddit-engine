package actors

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/proto"
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
    msg := context.Message()
    // print message type
    fmt.Printf("Received message: %T\n", msg)
    switch actorMsg := msg.(type) {
    case *actor.Started:
        fmt.Println("AuthActor received message")
    case *actor.Stopping:
        fmt.Println("AuthActor stopping")
    case *actor.Terminated:
        fmt.Println("AuthActor terminated")
	case *proto.RegisterRequest: 
        auth.RegisterNewUser(context, actorMsg)
    case *proto.LoginRequest:
        auth.LoginUser(context, actorMsg)
    case *proto.LogoutRequest:
        auth.LogoutUser(context, actorMsg)
    case *proto.TokenValidationRequest:
        auth.ValidateToken(context, actorMsg)
    default:
        fmt.Println("Unknown message")
	}
}

func (auth *AuthActor) RegisterNewUser(context actor.Context, actorMsg *proto.RegisterRequest) {
    fmt.Println("Registering new user")
    user, err := auth.service.RegisterNewUser(actorMsg.Username, actorMsg.Email, actorMsg.Password)
    response := &proto.RegisterResponse{
        Id:       0,
        Username: "",
        Error:    "",
    }
    
    if err != nil {
        fmt.Printf("Registration failed: %v\n", err)
        response.Error = err.Error()
    } else {
        response.Id = uint64(user.ID)
        response.Username = user.Username
    }
    
    context.Respond(response)
}

func (auth *AuthActor) LoginUser(context actor.Context, actorMsg *proto.LoginRequest) {
    token, err := auth.service.Login(actorMsg.Username, actorMsg.Password)
    if err != nil {
        context.Respond(&proto.LoginResponse{Error: err.Error()})
    }
    context.Respond(&proto.LoginResponse{Token: token})
}

func (auth *AuthActor) ValidateToken(context actor.Context, actorMsg *proto.TokenValidationRequest) {
    fmt.Println("Validating token")
    claims, err := auth.service.ValidateToken(actorMsg.Token)
    if err != nil {
        context.Respond(&proto.TokenValidationResponse{Error: err.Error()})
    } else {
        // jwt to proto claims
        claimsProto := &proto.Claims{
            UserId: uint64(claims.UserId),
            Username: claims.Username,
            // .. add more fields here
        }
        context.Respond(&proto.TokenValidationResponse{Valid: true, Claims: claimsProto})
    }
}

func (auth *AuthActor) LogoutUser(context actor.Context, actorMsg *proto.LogoutRequest) {
    err := auth.service.Logout(actorMsg.Token)
    if err != nil {
        context.Respond(&proto.LogoutResponse{Error: err.Error()})
    } else {
        context.Respond(&proto.LogoutResponse{})
    }
}