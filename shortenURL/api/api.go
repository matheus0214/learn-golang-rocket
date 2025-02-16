package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, r Response, s int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(r)
	if err != nil {
		slog.Error("Internal server error", "error", err)
		sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(s)
	if _, err = w.Write(data); err != nil {
		slog.Error("Error to send cliente response", "error", err)
		return
	}
}

const characters = "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)

	for i := range n {
		byts[i] = characters[rand.IntN(len(characters))]
	}

	return string(byts)
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(w, Response{Error: "Invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		_, err := url.Parse(body.URL)
		if err != nil {
			sendJSON(w, Response{Error: "Invalid url passed"}, http.StatusBadRequest)
			return
		}

		code := genCode()

		if _, ok := db[code]; ok {
			sendJSON(w, Response{Error: "URL already shortened"}, http.StatusConflict)
			return
		}

		db[code] = body.URL

		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		url, ok := db[code]
		if !ok {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
