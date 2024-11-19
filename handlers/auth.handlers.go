package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext, registrationActor *actor.PID) {
	var input messages.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send message to actor and wait for response
	future := rootContext.RequestFuture(registrationActor, &input, 5*time.Second)
	result, err := future.Result()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to register user: %v", err), http.StatusInternalServerError)
		return
	}

	registerResponse, ok := result.(*messages.RegisterResponse)
	if !ok {
		http.Error(w, "Invalid response from actor", http.StatusInternalServerError)
		return
	}

	if registerResponse.Error != "" {
		http.Error(w, registerResponse.Error, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(registerResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User registered. ID: %d, Username: %s\n", registerResponse.ID, registerResponse.Username)
}

// func fmtiPrintln(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext, authActorPid *actor.PID) {
// 	// validate user credentials
// 	// generate token
// }