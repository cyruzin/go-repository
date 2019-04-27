package model

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Movie type is a struct for movies.
type Movie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ReadRepository type is a interface to read from DB.
type ReadRepository interface {
	FindAll() []*Movie
	FindOneByID(id int64) *Movie
}

// WriteRepository type is a interface to write to DB.
type WriteRepository interface {
	Add(r *Movie) error
	Remove(id int64) error
	Update(id int64) error
}

// MovieRepository type is a struct for movies repository.
type MovieRepository struct {
	db *sqlx.DB
}

// NewMovieRepository initiate the service.
func NewMovieRepository(db *sqlx.DB) *MovieRepository {
	return &MovieRepository{db}
}

// FindAll func retrieves all movies.
func (r *MovieRepository) FindAll() ([]*Movie, error) {
	var movies []*Movie
	err := r.db.Select(&movies, `SELECT id,name FROM movie`)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return movies, nil
}

// FindOneByID func finds a movie by a given ID.
func (r *MovieRepository) FindOneByID(id int64) (*Movie, error) {
	var movie *Movie
	err := r.db.Get(&movie, `SELECT id,name FROM movie WHERE id = ?`, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return movie, nil
}

// Add func add a new movie.
func (r *MovieRepository) Add(m *Movie) error {
	if _, err := r.db.NamedExec(`INSERT INTO movie VALUES (:id, :name)`, m); err != nil {
		return err
	}
	return nil
}

// Update func updates the given movie.
func (r *MovieRepository) Update(m *Movie) error {
	if _, err := r.db.NamedExec("UPDATE movie SET name=:name WHERE id=:id", m); err != nil {
		return err
	}
	return nil
}

// Remove func removes a movie by a given ID.
func (r *MovieRepository) Remove(id int64) error {
	if _, err := r.db.Exec("DELETE FROM movie WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
