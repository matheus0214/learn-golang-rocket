package api

import (
	"gobid/internal/jsonutils"
	"net/http"

	"github.com/gorilla/csrf"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			_ = jsonutils.EncodeJSON(w, r, http.StatusUnauthorized, map[string]string{
				"error": "unauthorized access",
			})

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *Api) HandleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.EncodeJSON(w, r, http.StatusOK, map[string]string{
		"csrf_token": token,
	})
}
