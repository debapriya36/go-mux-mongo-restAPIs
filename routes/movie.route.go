package routes

import (
	"github.com/debapriya36/mongo-go-mux-crud/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/createMovie", controllers.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/v1/getMovies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/api/v1/updateMovie/{id}", controllers.UpdateWatchById).Methods("PUT")
	router.HandleFunc("/api/v1/deleteAllMovie", controllers.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/v1/getMovie/{name}", controllers.GetMovieByName).Methods("GET")
	return router
}
