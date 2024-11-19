package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/router"
	"github.com/nitingoyal0996/reddit-clone/actors"
	"github.com/nitingoyal0996/reddit-clone/database"
	"github.com/nitingoyal0996/reddit-clone/messages"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

func main() {
	// initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}
	defer sqlDB.Close()

	print("Database initialized.\n")

	system := actor.NewActorSystem()
	userRepo := repositories.NewUserRepository(db)
	authProducer := func() actor.Actor {
		return actors.NewAuthActor(userRepo, "chandukechacha")
	}
	authPoolProps := router.NewRoundRobinPool(5, actor.WithProducer(authProducer))
	pid := system.Root.Spawn(authPoolProps)

	registerFuture := system.Root.RequestFuture(pid, &messages.RegisterRequest{
		Username: "nitingoyal0996",
		Email: "nitin.goyal@gmail.com",
		Password: "password",
	}, 5*time.Second)

	// wait for the response
	registerResponse, err := registerFuture.Result()
	if err != nil {
		log.Fatalf("failed to register user: %v", err)
	} else {
		if registerResponse, ok := registerResponse.(*messages.RegisterResponse); ok {
			if registerResponse.Error != "" {
				log.Fatalf("Registration failed: %v", registerResponse.Error)
			} else {
				fmt.Printf("User registered. ID: %d, Username: %s\n", registerResponse.ID, registerResponse.Username)
			}
		}
	}
}
