package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"

	"./routes"
)

const (
	staticDir   = "./public"
	defaultPort = "8000"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/test", routes.TestApiHandler)
	r.HandleFunc("/api/check", routes.CheckApiHandler)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
	http.Handle("/", r)

	host := "0.0.0.0"
	port := getEnvOrDefault("PORT", defaultPort)
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

func getEnvOrDefault(env string, fallback string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}
	return fallback
}
