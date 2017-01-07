package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"./routes"
)

const (
	staticDir = "./public"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/test", routes.TestApiHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
	http.Handle("/", r)

	host := "0.0.0.0"
	port := "8000"
	listenAddr := host + ":" + port

	log.Printf("Server listening on %s", listenAddr)
	srv := &http.Server{
		Handler: r,
		Addr:    listenAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
