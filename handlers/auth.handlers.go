package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    future, err := h.Cluster.RequestFuture("auth", "Auth", &input)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to send request: %v", err), http.StatusInternalServerError)
        return
    }
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
    if json.NewEncoder(w).Encode(registerResponse); err != nil {
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

	// ID: login and Kind: Auth
	future, err := h.Cluster.RequestFuture("login", "Auth", &input)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to login user: %v", err), http.StatusInternalServerError)
		return
	}
	result, err := future.Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to login user: %v", err), http.StatusInternalServerError)
		return
	}

	loginResponse, ok := result.(*proto.LoginResponse)
	if !ok {
		http.Error(w, "Invalid response from actor", http.StatusInternalServerError)
		return
	}
	if loginResponse.Error != "" {
		http.Error(w, loginResponse.Error, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(loginResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
