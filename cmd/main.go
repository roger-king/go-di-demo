package main

import (
	"fx-demo/pkg/handlers"
	"fx-demo/pkg/logger"
	"fx-demo/pkg/server"
	"github.com/go-chi/chi"
	"go.uber.org/fx"
)

func register() {
	chi.NewMux()
}

func main() {
	app := fx.New(
		server.Module,
		logger.Module,
		fx.Provide(
			handlers.NewRouter,
			),
		)

	app.Run()
}
