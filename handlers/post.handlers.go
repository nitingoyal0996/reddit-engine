package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input messages.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	// create post
	payload := &messages.CreatePostRequest{
		PostSubID: input.PostSubID,
		PostTitle: input.PostTitle,
		PostBody: input.PostBody,
		PostOwner: input.PostOwner,
		PostKarma: input.PostKarma,
		PostCreationTime: input.PostCreationTime,
	}
	future, err = h.Cluster.RequestFuture("post", "Post", payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err = future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	CreatePostResponse, ok := result.(*messages.CreatePostResponse)
	if !ok {
		http.Error(w, "Invalid response from post actor", http.StatusInternalServerError)
		return
	}
	if CreatePostResponse.Error != "" {
		http.Error(w, CreatePostResponse.Error, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreatePostResponse)	
}

