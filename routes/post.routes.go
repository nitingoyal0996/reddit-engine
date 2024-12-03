// create posts subrouter and add routes to it

package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/handlers"
)

func PostRoutes(r *mux.Router, handler *handlers.Handler) {
	postRouter := r.PathPrefix("/posts").Subrouter()

	postRouter.HandleFunc("/", handler.CreatePostHandler).Methods("POST")
	postRouter.HandleFunc("/{id}", handler.GetPostHandler).Methods("GET")
	postRouter.HandleFunc("/{id}/comment", handler.CreateCommentHandler).Methods("POST")
	// postRouter.HandleFunc("/{id}/comments", handler.GetPostCommentsHandler).Methods("GET")
	postRouter.HandleFunc("{id}/vote", handler.UpdatePostVoteHandler).Methods("PUT")
}