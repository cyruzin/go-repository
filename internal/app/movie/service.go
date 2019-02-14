package movie

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Movie ...
type Movie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ReadRepository ...
type ReadRepository interface {
	FindAll() []*Movie
	FindOneByID(id int64) *Movie
}

// WriteRepository ...
type WriteRepository interface {
	Add(r *Movie) error
	Remove(id int64) error
	Update(id int64) error
}

// Repository ...
type Repository struct {
	db *sqlx.DB
}

// NewMovieRepository ...
func NewMovieRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}

// FindAll ...
func (r *Repository) FindAll() ([]*Movie, error) {
	var movies []*Movie

	err := r.db.Select(&movies, `SELECT * FROM movie`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return movies, nil
}

// FindOneByID ...
func (r *Repository) FindOneByID(id int64) (*Movie, error) {
	var movie *Movie

	err := r.db.Get(&movie, `SELECT * FROM movie WHERE id = $1`, id)

	return movie, err
}

// Add ...
func (r *Repository) Add(m *Movie) error {
	if _, err := r.db.NamedExec(`
		INSERT INTO movie VALUES (:id, :name)
	`, m); err != nil {
		return err
	}

	return nil
}

// Remove ...
func (r *Repository) Remove(id int64) error {
	if _, err := r.db.Exec("DELETE FROM movie WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}
