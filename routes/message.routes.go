package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/handlers"
)

func MessageRoutes(r *mux.Router, handler *handlers.Handler) {
    authRouter := r.PathPrefix("/messages").Subrouter()

    authRouter.HandleFunc("/create", handler.SendMessageHandler).Methods("POST")
    authRouter.HandleFunc("/", handler.GetMessagesHandler).Methods("GET")
}
