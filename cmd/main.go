package main

import (
	"fx-demo/pkg/handlers"
	"fx-demo/pkg/logger"
	"fx-demo/pkg/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		server.Module,
		logger.Module,
		handlers.Module,
	)

	app.Run()
}
