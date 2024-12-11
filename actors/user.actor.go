package actors

import (
	"fmt"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserActor struct {
	messageService *services.MessageService
}

func NewUserActor(msgRepo *repositories.SqliteMessageRepository) *UserActor {
	return &UserActor{
		messageService: services.NewMessageService(msgRepo),
	}
}

func (user *UserActor) Receive(context actor.Context) {
	msg := context.Message()
	fmt.Printf("UserActor Received message: %T\n", msg)
	switch actorMsg := msg.(type) {
	case *actor.Started:
		println("UserActor started")
	case *actor.Stopping:
		println("UserActor stopping")
	case *actor.Terminated:
		println("UserActor terminated")
	case *proto.SendMessageRequest:
		print("SendMessageRequest received\n")
		user.SendMessage(context, actorMsg)
	case *proto.GetMessagesRequest:
		print("GetMessagesRequest received\n")
		user.GetMessages(context, actorMsg)
	}
}

func (user *UserActor) SendMessage(context actor.Context, actorMsg *proto.SendMessageRequest) {
	// validate message
	authActor := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(authActor, &proto.TokenValidationRequest{Token: actorMsg.Token}, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.SendMessageResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.SendMessageResponse{Error: "Invalid token"})
	} else {
		fmt.Println("Token validated successfully")
		if err := user.messageService.SendMessage(actorMsg.Text, validationResponse.Claims.UserId, actorMsg.ToId); err != nil {
			fmt.Printf("%+v\n", err)
			context.Respond(&proto.SendMessageResponse{Error: err.Error()})
		} else {
			context.Respond(&proto.SendMessageResponse{Error: ""})
		}
	}
}

func (user *UserActor) GetMessages(context actor.Context, actorMsg *proto.GetMessagesRequest) {

	auth := cluster.GetCluster(context.ActorSystem()).Get("auth", "Auth")
	future := context.RequestFuture(auth, &proto.TokenValidationRequest{Token: actorMsg.Token}, 1*time.Second)
	res, err := future.Result()
	if err != nil {
		context.Respond(&proto.GetMessagesResponse{Error: err.Error()})
	}
	validationResponse, ok := res.(*proto.TokenValidationResponse)
	if !validationResponse.Valid || !ok {
		context.Respond(&proto.SendMessageResponse{Error: "Invalid token"})
	} else {
		print("\nToken validated successfully\n")
		if userMessages, err := user.messageService.GetMessages(validationResponse.Claims.UserId, actorMsg.ToId); err != nil {
			fmt.Printf("Error getting messages: %v\n", err)
			fmt.Printf("User ID: %v, To ID: %v\n", validationResponse.Claims.UserId, actorMsg.ToId)
			context.Respond(&proto.GetMessagesResponse{Error: err.Error()})
		} else {
			print("Messages retrieved successfully\n")
			protoMessages := make([]*proto.Message, len(userMessages))
			for i, msg := range userMessages {
				protoMessages[i] = &proto.Message{
					Id:        uint64(msg.ID),
					Text:      msg.Text,
					FromId:    uint64(msg.FromId),
					ToId:      uint64(msg.ToId),
					CreatedAt: timestamppb.New(msg.CreatedAt),
				}
			}
			context.Respond(&proto.GetMessagesResponse{Messages: protoMessages, Error: ""})
		}
	}
}