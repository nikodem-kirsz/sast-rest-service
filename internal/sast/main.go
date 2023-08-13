package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikodem-kirsz/sast-service/internal/common/server"
	"github.com/nikodem-kirsz/sast-service/internal/sast/ports"
)

func main() {
	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(), router)
	})
}
