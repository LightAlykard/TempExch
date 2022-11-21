package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Handlers ...
func (s *Service) Handlers() http.Handler {
	h := chi.NewMux()
	h.Route("/", func(r chi.Router) {
		h.Post("/login", s.login)
		h.Get("/logout", s.logout)
	})

	return h
}
