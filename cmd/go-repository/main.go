package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cyruzin/go-repository/internal/app/movie"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:@tcp(localhost:3306)/api-repository?parseTime=true")
	failOnError(err, "Failed to connect to DB")
	defer db.Close()

	r := movie.NewMovieRepository(db)

	// err = r.Add(&movie.Movie{
	// 	Name: "Batman",
	// })
	// failOnError(err, "Could not insert the movie")

	data, err := r.FindAll()
	failOnError(err, "Failed to fetch movies")
	j, err := json.MarshalIndent(data, "", "\t")
	failOnError(err, "Could not marshall movie response")
	fmt.Println(string(j))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("Error: %s. Message: %s\n", err, msg)
	}
}
