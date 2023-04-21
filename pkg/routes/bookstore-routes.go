package routes

import (
	"movielist/pkg/controller"

	"github.com/gorilla/mux"
)

var RegisterMoviestoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/movie/", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/", controller.GetMovie).Methods("GET")
	router.HandleFunc("/movie/{id}", controller.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{id}", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteMovie).Methods("DELETE")

}
