package main

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("CORS-HOST")},
		AllowCredentials: true,
		Debug:            true,
		AllowedMethods:   []string{"GET", "POST"},
	})

	http.ListenAndServe(":9090", c.Handler(HttpRouter().Init()))
}
