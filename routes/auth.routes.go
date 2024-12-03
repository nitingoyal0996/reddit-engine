package routes

import (
	"github.com/gorilla/mux"
	"github.com/nitingoyal0996/reddit-clone/handlers"
)

func AuthRoutes(r *mux.Router, handler *handlers.Handler) {
    authRouter := r.PathPrefix("/auth").Subrouter()

    authRouter.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
    authRouter.HandleFunc("/login", handler.LoginHandler).Methods("POST")
    authRouter.HandleFunc("/logout", handler.LogoutHandler).Methods("POST")
}
