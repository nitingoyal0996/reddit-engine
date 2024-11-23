package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/nitingoyal0996/reddit-clone/proto"
)

func (h *Handler) KarmaHandler(w http.ResponseWriter, r *http.Request, rootContext *actor.RootContext) {
	var input proto.KarmaRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	karmaActor := cluster.GetCluster(rootContext.ActorSystem()).Get("karma", "Karma")
	future := rootContext.RequestFuture(karmaActor, &input, 5*time.Second)
	res, err := future.Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	karmaResponse, ok := res.(*proto.KarmaResponse)
	if ok && karmaResponse.Error != "" {
		http.Error(w, karmaResponse.Error, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
