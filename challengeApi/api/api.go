package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	r *chi.Mux
	s *http.Server
}

func newServer() Server {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return Server{r: r, s: s}
}

func Serve() error {
	slog.Info("Starting server", "port", 8080)

	server := newServer()

	server.r.Route("/api", func(r chi.Router) {
		UserRoutes(r)
	})

	err := server.s.ListenAndServe()
	if err != nil {
		slog.Error("Error to initialize server", "error", err)
		return err
	}

	return nil
}
