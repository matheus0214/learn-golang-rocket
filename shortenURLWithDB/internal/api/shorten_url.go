package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"shortenURLWithDB/internal/store"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(db store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, apiResponse{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		code, err := db.SaveShortenedURL(r.Context(), body.URL)
		if err != nil {
			slog.Error("failed to create code", "error", err)
			sendJSON(w, apiResponse{Error: "internal server error"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
