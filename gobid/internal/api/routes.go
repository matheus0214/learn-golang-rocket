package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	// csrfMiddeware := csrf.Protect([]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), //dev only
	// )

	// api.Router.Use(csrfMiddeware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/csrftoken", api.HandleGetCSRFToken)

			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/login", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logout", api.handleLogoutUser)
			})

			r.Route("/products", func(r chi.Router) {
				r.Post("/", api.handleCreateProduct)
			})
		})
	})
}
