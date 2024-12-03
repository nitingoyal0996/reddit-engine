package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/handlers"
)

func SubredditRoutes(r *mux.Router, h *handlers.Handler) {
	SubredditRouter := r.PathPrefix("/subreddit").Subrouter()
	SubredditRouter.HandleFunc("/", h.CreateSubredditHandler).Methods("POST")
	SubredditRouter.HandleFunc("/{id}", h.GetSubredditHandler).Methods("GET")
	// SubredditRouter.HandleFunc("/{id}/posts", h.GetPostsBySubredditHandler).Methods("GET")
	SubredditRouter.HandleFunc("/{id}/subscribe", h.SubscribeSubredditHandler).Methods("POST")
	SubredditRouter.HandleFunc("/{id}/unsubscribe", h.UnsubscribeSubredditHandler).Methods("POST")

}