package main

import (
	"fmt"
	"log"
	"movielist/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterMoviestoreRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9000", r))
	fmt.Println("server was running")
}
