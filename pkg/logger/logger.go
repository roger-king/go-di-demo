package logger

import (
	"go.uber.org/fx"
	"log"
	"os"
)

var Module = fx.Options(
	fx.Provide(
		NewLogger,
		),
	)

type Logger interface {
	Println(v ...interface{})
}

func NewLogger() Logger {
	return log.New(os.Stdout, "[FX-DEMO] ", 0)
}
