package main

//  with out DB this rest api build

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "strconv"
	// "math/rand"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // Director pointer here pointing to Director struct below
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie // list of movies

func getMovies(w http.ResponseWriter, r *http.Request) { // http.Request should be pointer
	w.Header().Set("Content-Type", "application/json")
	// before sending we need to encode response w
	json.NewEncoder(w).Encode(movies) // sending process
}

func main() {
	var r = mux.NewRouter() // router by mux

	//sample movies added to list/struct
	movies = append(movies, Movie{ID: "1", Isbn: "23456", Title: "Movie One", Director: &Director{FirstName: "Jhon", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "34667", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "67342", Title: "Movie Three", Director: &Director{FirstName: "Devi", LastName: "DSP"}})

	// routers
	r.HandleFunc("movies", getMovies).Methods("GET")
	// r.HandleFunc("movies/{id}", getMovies).Methods("GET")
	// r.HandleFunc("movies", createMovie).Methods("POST")
	// r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
