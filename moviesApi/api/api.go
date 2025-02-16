package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shortenURL/omdb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie(apiKey))

	return r
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, r Response, s int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(r)
	if err != nil {
		sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(s)
	if _, err := w.Write(data); err != nil {
		slog.Error("Erro to send data to client", "error", err)
		return
	}
}

func handleSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")

		result, err := omdb.Search(apiKey, search)
		if err != nil {
			sendJSON(w, Response{Error: "Something wrong with ombd"}, http.StatusBadGateway)
			return
		}

		sendJSON(w, Response{Data: result}, http.StatusOK)
	}
}
