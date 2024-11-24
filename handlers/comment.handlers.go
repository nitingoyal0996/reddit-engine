package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func (h *Handler) CommentHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input messages.CreateCommentRequest
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

	// create comment
	payload := &messages.CreateCommentRequest{
		CommentID: input.CommentID,
		CommentBody: input.CommentBody,
		CommentPostID: input.CommentPostID,
		CommentOwner: input.CommentOwner,
		CommentKarma: input.CommentKarma,
		CommentParentID: input.CommentParentID,
		CommentCreationTime: input.CommentCreationTime,
	}
	future, err = h.Cluster.RequestFuture("comment", "comment", payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err = future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, ok := result.(*messages.CreateCommentResponse)
	if !ok {
		http.Error(w, "Invalid response from auth actor", http.StatusInternalServerError)
		return
	}
	if response.Error != "" {
		http.Error(w, response.Error, http.StatusInternalServerError)
		return
	}
}