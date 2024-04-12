package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movies struct {
	ID         string    `json: "id"`
	Isbn       string    `json:"isbn"`
	Title      string    `json:"title"`
	Director   *Director `json:"director"`
	Movie_year int       `json:"year"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// fixed number of thread worker
var num_worker = 10

// creating a channel, kind of like a queing system to take make sure that the goroutines are limited
var channel_worker = make(chan Movies)

// the function call to db to remove the movies
func workerDBCall() {
	for movies := range channel_worker {
		dbcall(movies.Movie_year, movies.ID)
	}

}

// in the subway scenario when we have multiple servers (workerdbcall), we are queuing the orders(request) each can take when
// iterating over the array, and executing them asynchronously.
// if there is not limit on how many threads can be created, we can use the function deleteMovieYear, and create
// threads over and over

func deleteMovieYearLimitedRoutines(year int) {
	// creating and starting the workers but are limited to num_workers, go routines will get limited here
	for i := 0; i < num_worker; i++ {
		go workerDBCall()
	}
	// iterating over the movies and sending to channel where it will get processed asynchronously by one the workers
	for index, movie := range movies {
		if movie.Movie_year < year {
			channel_worker <- movie
		}
	}

	close(channel_worker)
}

// This function will create routines for every index parsed, assumption unlimited CPU resources to use
func deleteMovieYear(year int) {
	for _, movies := range movies {
		if movies.Movie_year < year {
			// creating go routine to call to db
			go dbcall(year, movies.ID) // go routine running asynchronously not waiting for the db call api call with the year and movie id
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get the parameters from the URL
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovieYearRoutines(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	if err != nil {
		http.Error(w, "Invalid year parameter.", http.StatusBadRequest) // 400 code
		json.NewEncoder(w).Encode(err)
		return
	}
	go deleteMovieYear(year) // creates new go routine for deleting any movies before the given year, server is ready to accept new requests

}

var movies []Movies

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movies{ID: "1", Isbn: "43822", Title: "Movie One", Director: &Director{FirstName: "Francis Ford Coppola", LastName: "James Cameron"}})
	movies = append(movies, Movies{ID: "2", Isbn: "44833", Title: "Movie Two", Director: &Director{FirstName: "John Doe", LastName: "Margo"}})
	movies = append(movies, Movies{ID: "3", Isbn: "44586", Title: "Movie Three", Director: &Director{FirstName: "Manaswini Gogineni", LastName: "Flingo"}})
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/", deleteMovieYearRoutines).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
