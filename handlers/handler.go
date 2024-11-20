package handlers

import "github.com/asynkron/protoactor-go/cluster"

type Handler struct {
	Cluster *cluster.Cluster
}

func NewHandler(clusterProvider *cluster.Cluster) *Handler {
	return &Handler{
		Cluster: clusterProvider,
	}
}
