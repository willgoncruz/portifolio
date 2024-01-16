package main 

import (
	"log"
	"net/http"
	"os"

  "portifolio/delivery"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.HandleFunc("/", delivery.HandlePortifolio())

	log.Printf("connect to http://localhost:%s to see portifolio", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
