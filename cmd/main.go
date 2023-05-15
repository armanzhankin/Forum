package main

import (
	"log"
	"net/http"

	"forum/internals/handler"
	"forum/internals/repository"
	"forum/internals/service"
)

func main() {
	db, err := repository.InitRepo()
	if err != nil {
		log.Fatal("Database initialization failed: %w", err)
	}

	err = repository.CreateTables(db)
	if err != nil {
		log.Fatal("Creating tables failed: %w", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(*repo)
	handler := handler.NewHandler(services)

	s := &http.Server{
		Addr:    ":5050",
		Handler: handler.Routes(),
	}
	log.Printf("Starting server on http://localhost%s\n", s.Addr)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("%s", err.Error())
	}
}
