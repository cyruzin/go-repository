package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cyruzin/go-repository/internal/app/model"
)

// MovieService ...
type MovieService struct {
	r *model.MovieRepository
}

// NewMovieService ...
func NewMovieService(r *model.MovieRepository) *MovieService {
	return &MovieService{r}
}

// FindAll ...
func (mh MovieService) FindAll(w http.ResponseWriter, r *http.Request) {
	data, err := mh.r.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
