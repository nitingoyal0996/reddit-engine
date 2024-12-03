package handlers

import "github.com/asynkron/protoactor-go/actor"

type Handler struct {
	rootContext *actor.RootContext
}

func NewHandler(rootContext *actor.RootContext) *Handler {
	return &Handler{
		rootContext: rootContext,
	}
}
