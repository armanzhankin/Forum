package server

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Printf("Launching server on http://localhost%s\n", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	return err
}
