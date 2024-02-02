package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"tsis_1/internal"
)

var movies []internal.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&internal.Movie{})
}

// AddMovie adds a new movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie internal.Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = fmt.Sprintf("%d", len(movies)+1) // Simple ID generation
	movies = append(movies, newMovie)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newMovie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			var updatedMovie internal.Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]
			movies[index] = updatedMovie
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
	json.NewEncoder(w).Encode(&internal.Movie{})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
