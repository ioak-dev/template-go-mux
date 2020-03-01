package main

import (
	"github.com/ioak-dev/template-golang-service/src/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "8080"
	}

	r := mux.NewRouter()
	helloAPI := r.PathPrefix("/hello").Subrouter()
	helloAPI.HandleFunc("", controllers.GetHello).Methods(http.MethodGet)
	helloAPI.HandleFunc("", controllers.PostHello).Methods(http.MethodPost)
	helloAPI.HandleFunc("", controllers.PutHello).Methods(http.MethodPut)
	helloAPI.HandleFunc("", controllers.DeleteHello).Methods(http.MethodDelete)
	helloAPI.HandleFunc("", controllers.NotFoundHello)
	helloAPI.HandleFunc("/user/{userID}", controllers.ParamsHello).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":" + port, r))
}
