package ports

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nikodem-kirsz/sast-rest-service/internal/sast/app"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app}
}

func (h HttpServer) GetSastReports(w http.ResponseWriter, r *http.Request) {

}

func (h HttpServer) CreateSastReport(w http.ResponseWriter, r *http.Request) {

}

func (h HttpServer) DeleteReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}

func (h HttpServer) GetReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}

func (h HttpServer) UpdateReport(w http.ResponseWriter, r *http.Request, reportUUID uuid.UUID) {

}
