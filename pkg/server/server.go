package server

import (
	"context"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Options(
	fx.Invoke(RegisterHandlers, InitServer),
	)

func RegisterHandlers(handler http.Handler) {
	http.Handle("/", handler)
}

func InitServer(lc fx.Lifecycle) {
	server := &http.Server{
		Addr: ":8080",
	}

	lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Close()
		},
	})
}
