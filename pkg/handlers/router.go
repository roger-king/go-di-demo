package handlers

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(
		NewRouter,
		),
	)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})

	return router
}
