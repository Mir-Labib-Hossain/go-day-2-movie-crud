package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string   `json: "id"`
	Isbn     string   `json: "isbn"`
	Title    string   `json: "title"`
	Director Director `json: "director"`
}
type Director struct {
	FirstName string `json: "firstName"`
	LastName  string `json: "lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {}

func updateMovie(w http.ResponseWriter, r *http.Request) {}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	paramId := params["id"]
	for _, movie := range movies {
		if movie.Id == paramId {
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	paramsId := params["id"]
	for i, movie := range movies {
		if movie.Id == paramsId {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}

	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		Id: "0", Isbn: "34223", Title: "Life of pi", Director: Director{
			FirstName: "Labib",
			LastName:  "Hossain",
		},
	}, Movie{
		Id: "1", Isbn: "34229", Title: "Money Heist", Director: Director{
			FirstName: "Ayon",
			LastName:  "Chele",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8080", r))
}
