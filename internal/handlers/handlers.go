package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/meis1kqt/fastapiprac/internal/storage"
)

type Handlers struct {
	store *storage.TaskStore
}


func NewHandlers(store *storage.TaskStore) *Handlers {
	return &Handlers{store: store}
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}


func (h *Handlers) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks , err := h.store.GetAll()

	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON(w, 200, tasks)



}


func (h *Handlers) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, 405, "metod not allow")
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")


	id, err := strconv.Atoi(path)

	if err != nil {
		respondWithError(w, 400, "invalid id")
	}

	task , err := h.store.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, 404, "task not found")
			return
		}
		respondWithError(w, 500, err.Error())
		return
	}


	respondWithJSON(w, 200, task)

}