package handlers

import "github.com/meis1kqt/fastapiprac/internal/storage"

type Handlers struct {
	store *storage.TaskStore
}


func NewHandlers (store *storage.TaskStore) *Handlers {
	return &Handlers{store: store}
}