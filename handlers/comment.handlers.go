package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) CreateCommentHandler (w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.CreateCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	commentActor := cluster.GetCluster(rootContext.ActorSystem()).Get("comment", "Comment")
	future := rootContext.RequestFuture(commentActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	commentResponse, ok := res.(*proto.CreateCommentResponse)
	if ok && commentResponse.Error != "" {
		http.Error(w, commentResponse.Error, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	
}

func (h *Handler) GetCommentHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.GetCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	commentActor := cluster.GetCluster(rootContext.ActorSystem()).Get("comment", "Comment")
	future := rootContext.RequestFuture(commentActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentResponse, ok := res.(*proto.GetCommentResponse)
	if ok && commentResponse == nil {
		http.Error(w, "Failed.", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(commentResponse)
	}
	w.WriteHeader(http.StatusOK)
}