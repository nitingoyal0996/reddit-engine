package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/asynkron/protoactor-go/cluster/clusterproviders/automanaged"
	"github.com/asynkron/protoactor-go/cluster/identitylookup/disthash"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/nitingoyal0996/reddit-clone/actors"
	"github.com/nitingoyal0996/reddit-clone/database"
	"github.com/nitingoyal0996/reddit-clone/handlers"
	"github.com/nitingoyal0996/reddit-clone/repositories"
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

	// setup actor system
	authProps := actor.PropsFromProducer(func() actor.Actor {
        return actors.NewAuthActor(userRepo, "chanduKeChacha")
    })
	karmaProps := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewKarmaActor(userRepo)
	})
	authKind := cluster.NewKind("Auth", authProps)
	karmaKind := cluster.NewKind("Karma", karmaProps)

	// .. add more actor props here

	kinds := []*cluster.Kind{authKind, karmaKind}	// append more kinds here
	// Distributed hash lookup
	lookup := disthash.New()
	
    // New cluster definition
	config := remote.Configure("localhost", 8081)
	provider := automanaged.NewWithConfig(1*time.Second, 6331, "localhost:6331")
	clusterConfig := cluster.Configure("reddit-cluster", provider, lookup, config, cluster.WithKinds(kinds...))
	system := actor.NewActorSystem()
	cluster := cluster.New(system, clusterConfig)
    cluster.StartMember()
	// shutdown later
    defer cluster.Shutdown(true)

	rootContext := system.Root
	// declare HTTP handler to use cluster instead of actor system
    handler := handlers.NewHandler(cluster)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handler.RegisterHandler(w, r, rootContext)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handler.LoginHandler(w, r, rootContext)
	})

	http.HandleFunc("/user/karma", func(w http.ResponseWriter, r *http.Request) {
		handler.KarmaHandler(w, r, rootContext)
	})

	// .. add more handlers here

    http.ListenAndServe(":8080", nil)

	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
}