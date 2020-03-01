package main

import (
	"github.com/ioak-dev/template-golang-service/src/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	helloApi := r.PathPrefix("/hello").Subrouter()
	helloApi.HandleFunc("", controllers.GetHello).Methods(http.MethodGet)
	helloApi.HandleFunc("", controllers.PostHello).Methods(http.MethodPost)
	helloApi.HandleFunc("", controllers.PutHello).Methods(http.MethodPut)
	helloApi.HandleFunc("", controllers.DeleteHello).Methods(http.MethodDelete)
	helloApi.HandleFunc("", controllers.NotFoundHello)
	helloApi.HandleFunc("/user/{userID}", controllers.ParamsHello).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
