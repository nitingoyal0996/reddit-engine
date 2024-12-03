package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/handlers"
)

func UserRoutes(r *mux.Router, h *handlers.Handler)  {
	UserRouter := r.PathPrefix("/user").Subrouter()
	UserRouter.HandleFunc("/{id}/posts", h.GetPostsByUserHandler).Methods("GET")
	UserRouter.HandleFunc("/{id}/karma", h.KarmaHandler).Methods("PUT")
	// implement methods 
	// UserRouter.HandleFunc("/{id}/feed", h.GetFeedByUserHandler).Methods("GET")
	// UserRouter.HandleFunc("/{id}/comments", h.GetCommentsByUserHandler).Methods("GET")
}