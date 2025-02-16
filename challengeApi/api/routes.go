package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Get("/users", getUsers)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm alive"))
}
