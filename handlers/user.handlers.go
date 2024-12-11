package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) SendMessageHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userActor := cluster.GetCluster(rootContext.ActorSystem()).Get("user", "User")
	future := rootContext.RequestFuture(userActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, ok := res.(*proto.SendMessageResponse)
	if ok && response.Error != "" {
		http.Error(w, response.Error, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h* Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.GetMessagesRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userActor := cluster.GetCluster(rootContext.ActorSystem()).Get("user", "User")
	future := rootContext.RequestFuture(userActor, &input, 1*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, ok := res.(*proto.GetMessagesResponse)
	// print the response
	if ok && response.Error != "" {
		http.Error(w, response.Error, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.Messages)
}