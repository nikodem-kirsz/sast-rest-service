package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikodem-kirsz/sast-rest-service/internal/common/server"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/ports"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/service"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
