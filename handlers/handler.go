package handlers

import "github.com/brunompx/angula/storage"

type Handler struct {
	store *storage.Storage
}

func New(store *storage.Storage) *Handler {
	return &Handler{
		store: store,
	}
}
