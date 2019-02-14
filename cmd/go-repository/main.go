package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cyruzin/go-repository/internal/app/model"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/api-repository?parseTime=true")
	failOnError(err, "Failed to connect to DB")
	defer db.Close()

	mr := model.NewMovieRepository(db)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		data, err := mr.FindAll()
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8000", r)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("Error: %s. Message: %s\n", err, msg)
	}
}
