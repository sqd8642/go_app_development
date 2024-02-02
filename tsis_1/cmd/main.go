package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsis_1/internal"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetMovieByID).Methods("GET")
	router.HandleFunc("/movies", AddMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	initialMovies := []internal.Movie{
		{ID: "1", Title: "The Shawshank Redemption", Director: "Frank Darabont", Year: 1994},
		{ID: "2", Title: "The Godfather", Director: "Francis Ford Coppola", Year: 1972},
		{ID: "3", Title: "Pulp Fiction", Director: "Quentin Tarantino", Year: 1994},
		{ID: "4", Title: "Once Upon in Hollywood", Director: "Quentin Tarantino", Year: 2019},
		{ID: "5", Title: "The Hateful Eight", Director: "Quentin Tarantino", Year: 2015},
		{ID: "6", Title: "Reservoir Dogs", Director: "Quentin Tarantino", Year: 1992},
	}

	movies = append(movies, initialMovies...)

	port := ":8080"
	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
