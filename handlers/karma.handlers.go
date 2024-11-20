package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/nitingoyal0996/reddit-clone/messages"
)

func KarmaHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext, karmaActorPID *actor.PID) {
	var input messages.KarmaRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	future := rootContext.RequestFuture(karmaActorPID, &input, 5*time.Second)
	result, err := future.Result()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	karmaResponse, ok := result.(*messages.KarmaResponse)
	if !ok {
		http.Error(w, "Invalid response from actor", http.StatusInternalServerError)
		return
	}

	if karmaResponse.Error != "" {
		http.Error(w, karmaResponse.Error, http.StatusInternalServerError)
		return
	}
}