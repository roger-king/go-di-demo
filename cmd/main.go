package cmd

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
)

func register() {
	chi.NewMux()
}

func main() {
	app := fx.New(
		fx.Provide(),
		fx.Invoke(register),
		)

	app.Run()
}
