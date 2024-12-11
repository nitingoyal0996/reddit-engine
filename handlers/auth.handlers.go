package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authActor := cluster.GetCluster(rootContext.ActorSystem()).Get("auth", "Auth")
    future := rootContext.RequestFuture(authActor, &input, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting response: %v", err), http.StatusInternalServerError)
		return
	}

    registerResponse, ok := result.(*proto.RegisterResponse)
    if !ok {
        http.Error(w, "Invalid response type", http.StatusInternalServerError)
        return
    }
    if registerResponse.Error != "" {
        http.Error(w, registerResponse.Error, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(registerResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User registered. ID: %d, Username: %s\n", registerResponse.Id, registerResponse.Username)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authActor := cluster.GetCluster(rootContext.ActorSystem()).Get("auth", "Auth")
    future := rootContext.RequestFuture(authActor, &input, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting response: %v", err), http.StatusInternalServerError)
		return
	}
	loginResponse, ok := result.(*proto.LoginResponse)

	if !ok {
		http.Error(w, "Invalid response from actor", http.StatusInternalServerError)
		return
	}
	if loginResponse.Error != "" {
		if loginResponse.Error == "user not found" {
			http.Error(w, loginResponse.Error, http.StatusNotFound)
		} else if loginResponse.Error == "invalid password" {
			http.Error(w, loginResponse.Error, http.StatusUnauthorized)
		} else {
			http.Error(w, loginResponse.Error, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// logout handler
func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authActor := cluster.GetCluster(rootContext.ActorSystem()).Get("auth", "Auth")
	future := rootContext.RequestFuture(authActor, &input, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting response: %v", err), http.StatusInternalServerError)
		return
	}

	logoutResponse, ok := result.(*proto.LogoutResponse)
	if !ok {
		http.Error(w, "Invalid response from actor", http.StatusInternalServerError)
		return
	}
	if logoutResponse.Error != "" {
		http.Error(w, logoutResponse.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logoutResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}