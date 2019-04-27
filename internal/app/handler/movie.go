package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/go-repository/internal/app/model"
)

// MovieHandler type ...
type MovieHandler struct {
	r *model.MovieRepository
}

// NewMovie is a function that setups the handlers.
func NewMovie(r *model.MovieRepository) *MovieHandler {
	return &MovieHandler{r}
}

// FindAll lists the latest movies.
func (mh *MovieHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := mh.r.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
