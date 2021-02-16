package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest/controllers"
	"rest/di"
	"time"
)

func main() {
	di.BuildDIContainer()
	r := mux.NewRouter()
	r.HandleFunc("/create", controllers.CreateAdvert).Methods(http.MethodPost)
	r.HandleFunc("/adverts", controllers.ListAdverts).Methods(http.MethodGet)
	r.HandleFunc("/advert/{id:[0-9]+}", controllers.ViewAdvert).Methods(http.MethodGet)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8085",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
