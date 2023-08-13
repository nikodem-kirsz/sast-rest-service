package ports

import (
	"net/http"

	"github.com/google/uuid"
)

type HttpServer struct {
}

func NewHttpServer() HttpServer {
	return HttpServer{}
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
