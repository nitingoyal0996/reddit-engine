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
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/actors"
	"github.com/nitingoyal0996/reddit-clone/database"
	"github.com/nitingoyal0996/reddit-clone/handlers"
	"github.com/nitingoyal0996/reddit-clone/repositories"
	"github.com/nitingoyal0996/reddit-clone/routes"
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
	msgRepo := repositories.NewMessageRepository(db)
	subRepo := repositories.NewSubredditRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	// setup actor system
	authProps := actor.PropsFromProducer(func() actor.Actor {
        return actors.NewAuthActor(userRepo, "chanduKeChacha")
    })
	karmaProps := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewKarmaActor(userRepo)
	})
	userProps := actor.PropsFromProducer(func()actor.Actor {
		return actors.NewUserActor(msgRepo)
	})
	subProps := actor.PropsFromProducer(func()actor.Actor {
		return actors.NewSubredditActor(subRepo)
	})
	postProps := actor.PropsFromProducer(func()actor.Actor {
		return actors.NewPostActor(postRepo)
	})
	commentProps := actor.PropsFromProducer(func()actor.Actor {
		return actors.NewCommentActor(commentRepo)
	})
	// .. add more actor props here
	authKind := cluster.NewKind("Auth", authProps)
	karmaKind := cluster.NewKind("Karma", karmaProps)
	userKind := cluster.NewKind("User", userProps)
	subKind := cluster.NewKind("Subreddit", subProps)
	postKind := cluster.NewKind("Post", postProps)
	commentKind := cluster.NewKind("Comment", commentProps)
	// .. add more actor props here

	kinds := []*cluster.Kind{authKind, karmaKind, userKind, subKind, postKind, commentKind}	// append more kinds here
	// Distributed hash lookup
	lookup := disthash.New()
	
    // New cluster definition
	config := remote.Configure("127.0.0.1", 8080)
	provider := automanaged.NewWithConfig(1*time.Second, 6331, "localhost:6331")
	clusterConfig := cluster.Configure("reddit-cluster", provider, lookup, config, cluster.WithKinds(kinds...))
	system := actor.NewActorSystem()
	cluster := cluster.New(system, clusterConfig)
    cluster.StartMember()
	// shutdown later
    defer cluster.Shutdown(true)


	// initialize http server
	rootContext := system.Root
    handler := handlers.NewHandler(rootContext)
	router := mux.NewRouter()
	routes.AuthRoutes(router, handler)
	routes.UserRoutes(router, handler)
	routes.MessageRoutes(router, handler)
	routes.SubredditRoutes(router, handler)
	routes.PostRoutes(router, handler)
    http.ListenAndServe(":5678", router)

	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
}