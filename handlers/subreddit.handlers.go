package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func (h *Handler) SubredditHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input messages.GetSubredditRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//validate token
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

	// create subreddit
	payload := &messages.CreateSubredditRequest{
		SubName: input.SubName,
		SubDesc: input.SubDesc,
		SubCreationTime: input.SubCreationTime,
		SubOwner: input.SubOwner,
	}

	future, err = h.Cluster.RequestFuture("subreddit", "Subreddit", payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err = future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	CreateSubredditResponse, ok := result.(*messages.CreateSubredditResponse)
	if !ok {
		http.Error(w, "Invalid response from subreddit actor", http.StatusInternalServerError)
		return
	}
	if CreateSubredditResponse.Error != "" {
		http.Error(w, CreateSubredditResponse.Error, http.StatusInternalServerError)
		return
	}

}

