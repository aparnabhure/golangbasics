package router

import (
	"github.com/aparnabhure/golangmongodbapis/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
