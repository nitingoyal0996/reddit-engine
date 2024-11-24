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
	msgRepo := repositories.NewMessageRepository(db)
	subRepo := repositories.NewSubredditRepository(db)
	postRepo := repositories.NewPostRepository(db)
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
	// .. add more actor props here
	authKind := cluster.NewKind("Auth", authProps)
	karmaKind := cluster.NewKind("Karma", karmaProps)
	userKind := cluster.NewKind("User", userProps)
	subKind := cluster.NewKind("Subreddit", subProps)
	postKind := cluster.NewKind("Post", postProps)
	// .. add more actor props here

	kinds := []*cluster.Kind{authKind, karmaKind, userKind, subKind, postKind}	// append more kinds here
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
	http.HandleFunc("/user/messages", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Message request received: %s\n", r.Method)
		if r.Method == http.MethodGet {
			handler.GetMessagesHandler(w, r, rootContext)
		} else if r.Method == http.MethodPost {
			handler.SendMessageHandler(w, r, rootContext)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/user/subreddits", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateSubredditHandler(w, r, rootContext)
	})
	http.HandleFunc("/user/subreddits/subscribe", func(w http.ResponseWriter, r *http.Request) {
		handler.SubscribeSubredditHandler(w, r, rootContext)
	})
	http.HandleFunc("/user/subreddits/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		handler.UnsubscribeSubredditHandler(w, r, rootContext)
	})

	http.HandleFunc("/post/create", func(w http.ResponseWriter, r *http.Request) {
		handler.CreatePostHandler(w, r, rootContext)
	})
	http.HandleFunc("/post/get", func(w http.ResponseWriter, r *http.Request) {
		handler.GetPostHandler(w, r, rootContext)
	})
	http.HandleFunc("/post/get/user", func(w http.ResponseWriter, r *http.Request) {
		handler.GetPostsByUserHandler(w, r, rootContext)
	})
	http.HandleFunc("/post/get/subreddit", func(w http.ResponseWriter, r *http.Request) {
		handler.GetPostsBySubredditHandler(w, r, rootContext)
	})
	http.HandleFunc("/post/upvote", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdatePostVoteHandler(w, r, rootContext)
	})


    http.ListenAndServe(":5678", nil)

	// Run till a signal comes
	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
}