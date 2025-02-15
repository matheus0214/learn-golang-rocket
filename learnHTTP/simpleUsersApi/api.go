package simpleusersapi

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Password string

func (p Password) LogValue() slog.Value {
	return slog.StringValue("[REDACTED]")
}

type User struct {
	Username string   `json:"username"`
	ID       int64    `json:"id,string"`
	Role     string   `json:"role"`
	Password Password `json:"-"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Serve() {
	options := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, options))
	slog.SetDefault(l)

	slog.Info("Server initialize")

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {Username: "user1", ID: 1, Role: "admin", Password: "password"},
	}

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func sendJSON(w http.ResponseWriter, r Response, s int) {
	data, err := json.Marshal(r)
	if err != nil {
		slog.Error("Error to marshal: ", "error", err)
		sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(s)
	if _, err = w.Write(data); err != nil {
		slog.Error("Error to send response: ", "error", err)
		return
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		slog.Info("User", "user", user)
		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxError *http.MaxBytesError
			if errors.As(err, &maxError) {
				sendJSON(w, Response{Error: "Body to large"}, http.StatusRequestEntityTooLarge)
				return
			}

			slog.Error("Error to load user json", "error", err)
			sendJSON(w, Response{Error: "Internal server error"}, http.StatusInternalServerError)
			return
		}

		var user User
		if err = json.Unmarshal(data, &user); err != nil {
			sendJSON(w, Response{Error: "Invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		db[user.ID] = user
		w.WriteHeader(http.StatusCreated)
	}
}
