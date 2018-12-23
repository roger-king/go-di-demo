package handlers

import (
	"fx-demo/pkg/logger"
	"io"
	"net/http"
)

func NewRouter (logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println("Handler Called")
		io.WriteString(w, "Hello, FX-DEMO\n")
	})
}
