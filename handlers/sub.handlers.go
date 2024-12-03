package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) CreateSubredditHandler (w http.ResponseWriter, r *http.Request) {
	print("CreateSubredditHandler called")
	var input proto.CreateSubredditRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	subActor := cluster.GetCluster(h.rootContext.ActorSystem()).Get("subreddit", "Subreddit")
	future := h.rootContext.RequestFuture(subActor, &input, 5*time.Second)
	fmt.Printf("CreateSubredditHandler: %v\n", future)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	subResponse, ok := res.(*proto.CreateSubredditResponse)
	if ok && subResponse.Error != "" {
		http.Error(w, subResponse.Error, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	
}

// TODO: implement handler
func (h *Handler) GetSubredditHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) SubscribeSubredditHandler(w http.ResponseWriter, r *http.Request) {
	var input proto.SubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	subActor := cluster.GetCluster(h.rootContext.ActorSystem()).Get("subreddit", "Subreddit")
	future := h.rootContext.RequestFuture(subActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subResponse, ok := res.(*proto.SubscriptionResponse)
	if ok && !subResponse.Success {
		http.Error(w, subResponse.Message, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UnsubscribeSubredditHandler(w http.ResponseWriter, r *http.Request) {
	var input proto.UnsubscribeRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	subActor := cluster.GetCluster(h.rootContext.ActorSystem()).Get("subreddit", "Subreddit")
	future := h.rootContext.RequestFuture(subActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subResponse, ok := res.(*proto.UnsubscribeResponse)
	if ok && !subResponse.Success {
		http.Error(w, subResponse.Message, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}