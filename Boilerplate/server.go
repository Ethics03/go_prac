package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/joho/godotenv"

	"github.com/go-chi/cors"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()

	router.Mount("/v1", v1Router) // this nests another sub-router in the main router

	srv := &http.Server{

		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Port : ", portString)

	log.Printf("Listening to port : %v", portString)

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
