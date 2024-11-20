package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func (h *Handler) KarmaHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext, ) {
	var input messages.UpdateKarmaRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate request
	if input.Amount == 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// validate token
	future, err := h.Cluster.RequestFuture("auth", "Auth", &messages.TokenValidationRequest{Token: input.Token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	validationResponse, ok := result.(*messages.TokenValidationResponse)
	if !ok {
		http.Error(w, "Invalid response from auth actor", http.StatusInternalServerError)
		return
	}
	if !validationResponse.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// update karma
	payload := messages.KarmaRequest{
		UserId: validationResponse.Claims.UserId, 
		Amount: input.Amount,
	}
	future, err = h.Cluster.RequestFuture("karma", "Karma", &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err = future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	karmaResponse, ok := result.(*messages.KarmaResponse)
	if !ok {
		http.Error(w, "Invalid response from karma actor", http.StatusInternalServerError)
		return
	}
	if karmaResponse.Error != "" {
		http.Error(w, karmaResponse.Error, http.StatusInternalServerError)
		return
	}
}