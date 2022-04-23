package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"quality/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Get("/api/jobs", handlers.GetJobs)
	//run it on port 8080
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
