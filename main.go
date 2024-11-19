package main

import (
	"fmt"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/actors"
	"github.com/nitingoyal0996/reddit-clone/database"
	"github.com/nitingoyal0996/reddit-clone/handlers"
	"github.com/nitingoyal0996/reddit-clone/repositories"
)

var (
	system = actor.NewActorSystem()
    rootContext = system.Root
    
    registrationActorPID *actor.PID
	// .. more actors pid(s) may be added below
)

func main() {
	// initialize database
	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("failed to initialize database: %v\n", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("failed to get database connection: %v\n", err)
	}
	defer sqlDB.Close()
	fmt.Println("Database initialized.")

	// initialize data layer
	userRepo := repositories.NewUserRepository(db)
	authProps := actor.PropsFromProducer(func() actor.Actor {
		// jwt-secret = chanduKeChacha
        return actors.NewAuthActor(userRepo, "chanduKeChacha")
    })
    registrationActorPID = rootContext.Spawn(authProps)
	// ... more actors may be added below

	// handlers to serve actors with http requests
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterHandler(w, r, rootContext, registrationActorPID)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginHandler(w, r, rootContext, registrationActorPID)
	})
    http.ListenAndServe(":8080", nil)
}