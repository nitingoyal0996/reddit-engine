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


func (h *Handler) CreatePostHandler (w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	print("CreatePostHandler called")
	var input proto.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	subActor := cluster.GetCluster(rootContext.ActorSystem()).Get("post", "Post")
	future := rootContext.RequestFuture(subActor, &input, 5*time.Second)
	fmt.Printf("CreatePostHandler: %v\n", future)
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

func (h *Handler) GetPostHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.GetPostRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	postActor := cluster.GetCluster(rootContext.ActorSystem()).Get("post", "Post")
	future := rootContext.RequestFuture(postActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postResponse, ok := res.(*proto.GetPostResponse)
	if ok && postResponse == nil {
		http.Error(w, "Failed.", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(postResponse)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetPostsBySubredditHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.GetPostsBySubredditRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		print("GetPostsBySubredditHandler: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	postActor := cluster.GetCluster(rootContext.ActorSystem()).Get("post", "Post")
	future := rootContext.RequestFuture(postActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postResponse, ok := res.(*proto.GetPostsBySubredditResponse)
	if ok && postResponse == nil {
		http.Error(w, "Failed to fetch posts.", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(postResponse)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetPostsByUserHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.GetPostByUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	postActor := cluster.GetCluster(rootContext.ActorSystem()).Get("post", "Post")
	future := rootContext.RequestFuture(postActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postResponse, ok := res.(*proto.GetPostByUserResponse)
	if ok && postResponse == nil {
		http.Error(w, "Post creation failed.", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(postResponse)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdatePostVoteHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.UpdatePostVoteRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	postActor := cluster.GetCluster(rootContext.ActorSystem()).Get("post", "Post")
	future := rootContext.RequestFuture(postActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postResponse, ok := res.(*proto.UpdatePostVoteResponse)
	if ok && postResponse == nil {
		http.Error(w, "Post creation failed.", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(postResponse)
	}
	w.WriteHeader(http.StatusOK)
}