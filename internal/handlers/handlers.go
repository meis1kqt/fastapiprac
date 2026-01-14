package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/meis1kqt/fastapiprac/internal/storage"
)

type Handlers struct {
	store *storage.TaskStore
}


func NewHandlers(store *storage.TaskStore) *Handlers {
	return &Handlers{store: store}
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "applicatiin/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}


func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks , err := h.store.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve tasks")
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}	

func (h *Handlers) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	task , err := h.store.GetById()

	if err != nil {
		log.Fatalf("failed")
	}

	respondWithJSON(w, http.StatusOK, task)
}

