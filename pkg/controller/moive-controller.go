package controller

import (
	"encoding/json"
	"fmt"
	"movielist/pkg/models"
	"movielist/pkg/utils"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "gorm.io/gorm/utils"
)

var NewMovie models.Movie

func GetMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := models.GetMovie()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["id"]
	Id, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	movieDetails, _ := models.GetMovieById(Id)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	createMovies := &models.Movie{}
	utils.ParseBody(r, createMovies)
	b := createMovies.CreateMovie()
	res, _ := json.Marshal(b)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	moviesId := vars["id"]
	Id, err := strconv.ParseInt(moviesId, 0, 0)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	movie := models.DeleteMovie(Id)
	res, _ := json.Marshal(movie)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &models.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["id"]
	Id, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	movieDetails, db := models.GetMovieById(Id)
	if updateMovie.Title != "" {
		movieDetails.Title = updateMovie.Title
	}
	if updateMovie.Rating != "" {
		movieDetails.Rating = updateMovie.Rating
	}
	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
