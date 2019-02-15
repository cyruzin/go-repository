package main

import (
	"log"
	"net/http"

	"github.com/cyruzin/go-repository/internal/app/handler"
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
	ms := handler.NewMovie(mr)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/movies", ms.FindAll)

	http.ListenAndServe(":8000", r)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("Error: %s. Message: %s\n", err, msg)
	}
}
