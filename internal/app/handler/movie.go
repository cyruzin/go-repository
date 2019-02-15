package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/go-repository/internal/app/model"
)

// MovieHandler ...
type MovieHandler struct {
	r *model.MovieRepository
}

// NewMovieHandler ...
func NewMovieHandler(r *model.MovieRepository) *MovieHandler {
	return &MovieHandler{r}
}

// FindAll ...
func (mh MovieHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := mh.r.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
