package api

import (
	"errors"
	"log/slog"
	"net/http"
	"shortenURLWithDB/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenURL(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		fullURL, err := db.GeFullURL(r.Context(), code)
		if err != nil {
			slog.Error("failed to get code", "error", err)
			if errors.Is(err, redis.Nil) {
				sendJSON(w, apiResponse{Error: "code not found"}, http.StatusNotFound)
				return
			}

			sendJSON(w, apiResponse{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: getShortenedURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
