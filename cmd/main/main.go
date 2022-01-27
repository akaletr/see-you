package main

import (
	"cmd/main/main.go/internal/place"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	placeHandler := place.NewHandler()
	placeHandler.Register(router)

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: router,
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
